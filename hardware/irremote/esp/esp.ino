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
  int frequency;
  unsigned int rawData;
};

void setup() {
  Serial.begin(9600);
  SPI.begin();

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
  client.setServer(mqtt_server, mqtt_port);
  client.setCallback(handler);
}

void loop() {
  if (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    if (client.connect(device_id, "test", "test")) {
      Serial.println("connected");
      client.subscribe(getTopicWildcard().c_str());
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

  String jsonStr = "";
  unsigned int i = 0;
  while (i < length)
    jsonStr += (char)p[i++];
  //

  String _topicStr = String(topic);
  String prefix = String("huberlink/") + String(device_id);
  if (!_topicStr.startsWith(prefix)) {
    Serial.println("Topic not match");
    return;
  }
  // split topic from huberlink/device_id/topicname
  String topicName = _topicStr.substring(prefix.length() + 1);
  Serial.println(topicName);
  // do exuecute
  DynamicJsonDocument doc(1024);
  deserializeJson(doc, jsonStr);
  Command cmd;
  cmd.frequency = doc["frequency"];
  cmd.rawData = doc["rawData"];
  onExecuteCommand(&cmd);
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
  Serial.println(cmd->frequency);
  Serial.println(cmd->rawData);
  free(cmd);
  // sent raw data to arduino uno
  // then report status whic success or not
}