#include "WString.h"
const char* device_id = "1638187483118178304";
const char* device_token = "eOq830I6_QRpE3W-S1F_hx1Zj9BGczeVOLf4ieW9J09KWIHq5AElaMxm3KhFx6bt8zSOjJwUUCW9TQvml4wTI7g2PiBX8mdaNCHdA97JrREoCIvYN8BSnDHVMCJB6bnfv4iMKxte8OehrpyJehC9TKN1gt774w6_yRMsOUVU0-HAohyoJ5ICfWEGngkhEW5HnQwbmybdnuPVhYSgZ-z8uUqQpSE2UXMuUAZ0Z-5_u3R0bgcwyP9Z2SjHcndvxvGFQPCoaq8pRi3wMsFMaZvwfqNMY86D17t1AL7cTJ2uqJWwsROU4G0wp2rr3ciZZVMo0yXdb4PT4DqCvEjtFlYEMH57eaWUN-lyw6WRMnqzBSLwuFbtvRDFJ7HAvyrxbv_9-zm8w_7n3X-tcvasIVchAxXjwkzOAJZe4qy7iYwxRxEps5rO3JENllWlS0bR5LpYir8veiAT5Yyp7Ji6xIv-cecV7S9mJ6megLTYeeh5g4cky5B1IL6MflRwNOsNsNAk2nFQ8mnqL2hZXlQQvVAWkz-8KHsZG9AGpiiOK6qnzYJMDqVJ7VsG_QwmjPCz2yk4a2wasKXnPCcIor0rzzlCQoQLLcNkv_sOofPPM9gK1aZ9_s1RluVxWT8eNdBjW4acN6MXUWjyYbOZmYuKBOgQMMGNsj4=";

const char* topic_heartbeat = "/thing/heartbeat";
const char* topic_execute = "/thing/irremote/execute";
const char* topic_learning = "/thing/irremote/learning";
const char* topic_learning_result = "/thing/irremote/learning/result";
const char* topic_ping = "/thing/ping";
String getTopicLearningResult() {
  return String("huberlink/") + String(device_id) + String(topic_learning_result);
}
String getHeartbeatTopic() {
  return String("huberlink/") + String(device_id) + String(topic_heartbeat);
}
String getExecuteTopic() {
  return String("huberlink/") + String(device_id) + String(topic_execute);
}
String getLearningTopic() {
  return String("huberlink/") + String(device_id) + String(topic_learning);
}
String getPingTopic() {
  return String("huberlink/") + String(device_id) + String(topic_ping);
}
const int MAX_RAW_DATA = 10;


#define mqtt_server "141.11.156.252"
#define mqtt_port 1883
#define mqtt_user "huberlink"
#define mqtt_password "huberlink"

char* fmtString(const char* format, const char* string) {
  size_t len = snprintf(NULL, 0, format, string) + 1;
  char* buf = new char[len];
  snprintf(buf, len, format, string);
  return buf;
}