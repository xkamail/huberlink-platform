#include <ArduinoJson.h>
#include <WiFiManager.h>
#include <PubSubClient.h>
#include <SoftwareSerial.h>
#include <string.h>
#include "config.h"
#include <SPI.h>

WiFiManager wm;
WiFiClient espClient;
PubSubClient client(espClient);

// Command that sent from mqtt to run the remote code
struct Command {
  unsigned int rawData[400] = {};
};

void setup() {
  pinMode(D1, INPUT);
  Serial.begin(115200);
  //attachInterrupt(digitalPinToInterrupt(D0), dataFromUno, RISING);

  SPI.begin();
  pinMode(SS, OUTPUT);
  // setup SPI
  digitalWrite(SS, HIGH);
  // automatically connect using saved credentials if they exist
  // If connection fails it starts an access point with the specified name
  if (wm.autoConnect("IRRemote HuberLink")) {
    Serial.println("connected...yeey :)");
  } else {
    Serial.println("Configportal running");
  }

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  client.setBufferSize(500);
  client.setServer(mqtt_server, mqtt_port);
  client.setCallback(handler);
}

void loop() {
  if (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    if (client.connect(device_id, mqtt_user, mqtt_password)) {
      Serial.println("connected");
      client.subscribe(getLearningTopic().c_str());
      client.subscribe(getExecuteTopic().c_str());
      client.subscribe(getPingTopic().c_str());
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      delay(5000);
      return;
    }
    // fix some issue wifi stability
    delay(10);
  }

  // process logic
  onHeartbeat();
  client.loop();
  wm.process();
}


void handler(char *topic, byte *p, unsigned int length) {

  String _topicStr = String(topic);
  String prefix = String("huberlink/") + String(device_id);
  if (!_topicStr.startsWith(prefix)) {
    Serial.println("Topic not match");
    return;
  }
  String topicName = _topicStr.substring(prefix.length() + 1);
  Serial.print("TOPIC:");
  Serial.println(topicName);

  if (topicName == "thing/irremote/execute") {
    Serial.print("Size:");
    Serial.println(length);
    char *token = strtok((char *)p, ",");
    int _i = 0;
    uint8_t codes[800] = {};
    while (token != NULL) {
      uint8_t n = (uint8_t)atoi(token);
      codes[_i++] = n;
      token = strtok(NULL, ",");
    }
    digitalWrite(SS, LOW);
    //
    for (int i = 0; i < _i; i++) {
      SPI.transfer(codes[i]);
      Serial.print(codes[i]);
      Serial.print(" ");
    }
    SPI.transfer(0);  // end
    Serial.println();
    //
    digitalWrite(SS, HIGH);
    return;
  }
  if (topicName == "thing/ping") {
    return;
  }
}

unsigned long latestBeat = 0;
void onHeartbeat() {
  // check if mqtt is connected and wifi also
  unsigned long currentMillis = millis();
  if (currentMillis - latestBeat >= 5000) {
    latestBeat = currentMillis;
    char *payload = fmtString("hello %s", String("world").c_str());
    Publish(getHeartbeatTopic().c_str(), payload);
  }
}

void Publish(const char *topic, const char *payload) {
  // trim /
  while (topic[0] == '/')
    topic++;
  client.publish(topic, payload);
}

void onExecuteCommand(Command *cmd) {
  //
  Serial.println("Execute command");
  free(cmd);
  // sent raw data to arduino uno
  // then report status whic success or not
}
