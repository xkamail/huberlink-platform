#include <ArduinoJson.h>
#include <WiFiManager.h>
#include "PinAndDefinations.h"
#include <SoftwareSerial.h>
#include <PubSubClient.h>
#include <string.h>

#define mqtt_server "191.101.214.207"
#define mqtt_port 1883
#define mqtt_user "test"
#define mqtt_password "test"

#define DEVICE_TOKEN "eOq830I6_QRpE3W-S1F_hx1Zj9BGczeVOLf4ieW9J09KWIHq5AElaMxm3KhFx6bt8zSOjJwUUCW9TQvml4wTI7g2PiBX8mdaNCHdA97JrREoCIvYN8BSnDHVMCJB6bnfv4iMKxte8OehrpyJehC9TKN1gt774w6_yRMsOUVU0-HAohyoJ5ICfWEGngkhEW5HnQwbmybdnuPVhYSgZ-z8uUqQpSE2UXMuUAZ0Z-5_u3R0bgcwyP9Z2SjHcndvxvGFQPCoaq8pRi3wMsFMaZvwfqNMY86D17t1AL7cTJ2uqJWwsROU4G0wp2rr3ciZZVMo0yXdb4PT4DqCvEjtFlYEMH57eaWUN-lyw6WRMnqzBSLwuFbtvRDFJ7HAvyrxbv_9-zm8w_7n3X-tcvasIVchAxXjwkzOAJZe4qy7iYwxRxEps5rO3JENllWlS0bR5LpYir8veiAT5Yyp7Ji6xIv-cecV7S9mJ6megLTYeeh5g4cky5B1IL6MflRwNOsNsNAk2nFQ8mnqL2hZXlQQvVAWkz-8KHsZG9AGpiiOK6qnzYJMDqVJ7VsG_QwmjPCz2yk4a2wasKXnPCcIor0rzzlCQoQLLcNkv_sOofPPM9gK1aZ9_s1RluVxWT8eNdBjW4acN6MXUWjyYbOZmYuKBOgQMMGNsj4="

const char *DEVICE_ID = "1633177370523340800";
// what wrong of this line ?
const char topicConfig[100], topicState[100], topicExecute[100];
sprintf(topic, "devices/%s", DEVICE_ID);
sprintf(topicConfig, "devices/%s", DEVICE_ID);
sprintf(topicConfig, "devices/%s", DEVICE_ID);



WiFiManager wm;
WiFiClient espClient;
SoftwareSerial uno(RX, TX);
PubSubClient client(espClient);

const byte numChars = 32;
char receivedChars[numChars];
boolean newData = false;

void setup() {
  // put your setup code here, to run once:
  Serial.begin(115200);

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
  //
  client.setServer(mqtt_server, mqtt_port);
  client.setCallback(MQTTHandler);
}

void loop() {
  if (!client.connected()) {
    Serial.print("Attempting MQTT connection...");
    if (client.connect(DEVICE_ID, "test", "test")) {
      Serial.println("connected");
      client.subscribe("/ESP/LED");
    } else {
      Serial.print("failed, rc=");
      Serial.print(client.state());
      Serial.println(" try again in 5 seconds");
      delay(5000);
      return;
    }
  }

  // process logic
  onHeartbeat();
  client.loop();
  wm.process();
  //
  serialEvent();
  sendNewData();
}

void serialEvent() {
  static boolean recvInProgress = false;
  static byte ndx = 0;
  char startMarker = '<';
  char endMarker = '>';
  char rc;

  while (Serial.available() > 0 && newData == false) {
    rc = Serial.read();

    if (recvInProgress == true) {
      if (rc != endMarker) {
        receivedChars[ndx] = rc;
        ndx++;
        if (ndx >= numChars) {
          ndx = numChars - 1;
        }
      } else {
        receivedChars[ndx] = '\0';  // terminate the string
        recvInProgress = false;
        ndx = 0;
        newData = true;
      }
    }

    else if (rc == startMarker) {
      recvInProgress = true;
    }
  }
}
void sendNewData() {
  //
  if (newData == true) {
    Serial.print("This just in ... ");
    Serial.println(receivedChars);
    newData = false;
  }
}

void MQTTHandler(char *topic, byte *p, unsigned int length) {
  Serial.print("Message: ");
  Serial.print(topic);
  Serial.prinln("");

  String jsonStr = "";
  unsigned int i = 0;
  while (i < length)
    jsonStr += (char)p[i++];
  //

  DynamicJsonDocument doc(1024);
  deserializeJson(doc, json);

  const char *sensor = doc["sensor"];
  long time = doc["time"];
  double latitude = doc["data"][0];
  double longitude = doc["data"][1];
  // check if topic start with /devices/{device_id}
  // then check if topic is /devices/{device_id}/command
  // then check if topic is /devices/{device_id}/config
  // then check if topic is /devices/{device_id}/state

  // check string has prefix devices/{device_id} or not
  String topicStr = String(topic);
  String prefix = String("devices/") + String(DEVICE_ID);
  if (!topicStr.startsWith(prefix)) {
    Serial.println("Topic not match");
    return;
  }
}

unsigned long latestBeat = 0;
void onHeartbeat() {
  // check if mqtt is connected and wifi also
  // then publish a topic into mqtt server
  // topic: /generic/{device_id}/heartbeat
  unsigned long currentMillis = millis();
  if (currentMillis - latestBeat >= 5000) {
    latestBeat = currentMillis;
    Serial.println("beat~");
    char str[50];
    sprintf(str, "generic/%s/heartbeat", String(DEVICE_ID).c_str());
    Publish(str, "Hi");
  }
}

void Publish(const char *topic, const char *payload) {
  // trim /
  while (topic[0] == '/')
    topic++;
  client.publish(topic, payload);
}
