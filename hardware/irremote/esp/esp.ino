#define MARK_EXCESS_MICROS 20
#define RAW_BUFFER_LENGTH 600
#include <WiFiManager.h>
#include <PubSubClient.h>
#include <string.h>
#include "config.h"
#include <SPI.h>
#include <IRremote.hpp>

WiFiManager wm;
WiFiClient espClient;
PubSubClient client(espClient);

void setup() {
  pinMode(D1, INPUT);
  Serial.begin(115200);
  //attachInterrupt(digitalPinToInterrupt(D0), dataFromUno, RISING);
  IrReceiver.begin(D4, false);
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
  listeningIR();
}
volatile bool feedbackResult = false;

void listeningIR() {
  if (IrReceiver.decode() && !feedbackResult) {
    if (IrReceiver.decodedIRData.rawDataPtr->rawlen < 4) {
      Serial.print(F("Ignore data with rawlen="));
      Serial.println(IrReceiver.decodedIRData.rawDataPtr->rawlen);
      return;
    }
    if (IrReceiver.decodedIRData.flags & IRDATA_FLAGS_IS_REPEAT) {
      Serial.println(F("Ignore repeat"));
      return;
    }
    if (IrReceiver.decodedIRData.flags & IRDATA_FLAGS_IS_AUTO_REPEAT) {
      Serial.println(F("Ignore autorepeat"));
      return;
    }
    if (IrReceiver.decodedIRData.flags & IRDATA_FLAGS_PARITY_FAILED) {
      Serial.println(F("Ignore parity error"));
      return;
    }
    if (IrReceiver.decodedIRData.flags & IRDATA_FLAGS_WAS_OVERFLOW) {
      return;
    }
    uint8_t raw[600] = {};
    int rawLength = IrReceiver.decodedIRData.rawDataPtr->rawlen - 1;
    IrReceiver.compensateAndStoreIRResultInArray(raw);
    String result;
    for (int i = 0; i < rawLength; i++) {
      result += String(raw[i]) + String(",");
    }
    feedbackResult = true;
    Publish(getTopicLearningResult().c_str(), result.c_str());
  }
}

void handler(char *topic, byte *p, unsigned int length) {

  String _topicStr = String(topic);
  String prefix = String("huberlink/") + String(device_id);
  if (!_topicStr.startsWith(prefix)) {
    Serial.println("Topic not match");
    return;
  }
  String topicName = "/" + _topicStr.substring(prefix.length() + 1);
  Serial.print("TOPIC:");
  Serial.println(topicName);
  if (topicName == String(topic_execute)) {
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
    }
    SPI.transfer(0);  // end
    Serial.println();
    //
    digitalWrite(SS, HIGH);
    return;
  }
  if (topicName == String(topic_learning)) {
    feedbackResult = false;
    Serial.println("start learning");
    IrReceiver.start();
    return;
  }
  if (topicName == String(topic_ping)) {
    if (length == 4) {  // prevent self message
                        // TODO: improve check message contain
                        // instead of length of string
      return;
    }
    Publish(getPingTopic().c_str(), String("pong").c_str());
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
