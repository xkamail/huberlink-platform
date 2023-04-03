# huberlink-platform

HuberLink Platform IoT

## MQTT Topic

`huberlink/%s/thing/execute` to execute command

`huberlink/%s/thing/heartbeat` to check latest state

`huberlink/%s/thing/report` to save data which come from thing

```mermaid
sequenceDiagram
  autonumber
  participant W as Web Application
  participant A as Web Service API
  participant D as Database
  participant MQ as MQTT Broker
  participant M as Microcontroller

  W-->>A: ต้องการกดปุ่ม
  A-->>D: query หา code ของปุ่มนั้น
  D->>A: ตอบกลับค่า code ที่หาเจอ
  A->>MQ: publish ข้อมูล code ที่เป็น array เข้าไป

  MQ->>M: ได้รับข้อมูล code infrared uint8_t[]
  M->M: ดำเนินการส่งสัญญาณออกไป LED

```
