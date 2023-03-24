
#include <IRLibRecvPCI.h>
#include <IRLibSendBase.h>  //We need the base code
#include <IRLib_HashRaw.h>
#define IR_RECEIVE_PIN 2  // To be compatible with interrupt example, pin 2 is chosen here.
#define IR_SEND_PIN 3
#define TONE_PIN 4
#define APPLICATION_PIN 5
#define ALTERNATIVE_IR_FEEDBACK_LED_PIN 6  // E.g. used for examples which use LED_BUILDIN for example output.
#define _IR_TIMING_TEST_PIN 7

#include "avr/interrupt.h"

IRsendRaw sender;

IRrecvPCI receiver(2);  // D2
#if !defined(STR_HELPER)
#define STR_HELPER(x) #x
#define STR(x) STR_HELPER(x)
#endif


void setup() {

  pinMode(4, OUTPUT);
  pinMode(2, INPUT);
  //
  while (!Serial) {}
  Serial.begin(9600);
  Serial.println("start");
  receiver.enableIRIn();

  receiver.setFrameTimeout(100000);
}

void loop() {
  sw();
  handleResult();
}

void handleResult() {
  if (receiver.getResults()) {
    Serial.println(F("Got a result"));
    Serial.println(recvGlobal.recvLength, DEC);
    Serial.print(F("uint16_t rawData[RAW_DATA_LEN]={\n\t"));
    for (bufIndex_t i = 1; i < recvGlobal.recvLength; i++) {
      Serial.print(recvGlobal.recvBuffer[i], DEC);
      Serial.print(F(", "));
      if ((i % 8) == 0) Serial.print(F("\n\t"));
    }
    Serial.println(F("1000};"));  //Add arbitrary trailing space
    // receiver.enableIRIn();        //Restart receiver
  }
}
//recvGlobal
void sent() {
  Serial.println("sent");
  if (recvGlobal.recvBuffer[recvGlobal.recvLength - 1] != 1000) {
    Serial.println("[DEBUG] rearrange array!");
    for (bufIndex_t i = 0; i < recvGlobal.recvLength; i++) {
      recvGlobal.recvBuffer[i] = recvGlobal.recvBuffer[i + 1];
    }
    recvGlobal.recvBuffer[recvGlobal.recvLength - 1] = 1000;
  }
  Serial.println("finished");
  for (bufIndex_t i = 0; i < recvGlobal.recvLength; i++) {
    Serial.print(recvGlobal.recvBuffer[i], DEC);
    Serial.print(F(", "));
    if ((i % 8) == 0) Serial.print(F("\n\t"));
  }

  sender.send(recvGlobal.recvBuffer, recvGlobal.recvLength, 38);
  //
  //receiver.enableIRIn();
}

long debounce = 0;

void sw() {
  int sw = digitalRead(5);
  if ((millis() - debounce) > 1000) {
    if (sw == LOW) {
      Serial.println("press sw");
      sent();
      debounce = millis();
    }
  }

  return;
}
ISR(TIMER1_OVF_vect) {
  PORTB ^= _BV(5);  // for debug
}
