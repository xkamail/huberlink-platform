
#include <IRLibRecvPCI.h>
#include <IRLib_HashRaw.h>
#define IR_RECEIVE_PIN 2  // To be compatible with interrupt example, pin 2 is chosen here.
#define IR_SEND_PIN 3
#define TONE_PIN 4
#define APPLICATION_PIN 5
#define ALTERNATIVE_IR_FEEDBACK_LED_PIN 6  // E.g. used for examples which use LED_BUILDIN for example output.
#define _IR_TIMING_TEST_PIN 7
#include <IRremote.hpp>

#include "avr/interrupt.h"
IRrecvPCI receiver(2);  // D2
#if !defined(STR_HELPER)
#define STR_HELPER(x) #x
#define STR(x) STR_HELPER(x)
#endif

IRsend IrSender;

void setup() {

  pinMode(4, OUTPUT);
  pinMode(2, INPUT);
  //
  while (!Serial) {}
  Serial.begin(9600);
  Serial.println("start");
  receiver.enableIRIn();

  receiver.setFrameTimeout(100000);
  Serial.println(F("START " __FILE__ " from " __DATE__ "\r\nUsing library version " VERSION_IRREMOTE));
  Serial.println(F("Send IR signals at pin " STR(IR_SEND_PIN)));

  IrSender.begin();  // Start with IR_SEND_PIN as send pin and enable feedback LED at default feedback LED pin
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
    receiver.enableIRIn();        //Restart receiver
  }
}
uint16_t rawData[68] = {
  4474, 4602, 474, 1774, 474, 1774, 474, 1774,
  474, 650, 478, 646, 478, 646, 454, 670,
  478, 646, 478, 1774, 474, 1774, 474, 1774,
  458, 666, 454, 670, 454, 670, 454, 670,
  478, 650, 474, 650, 474, 1774, 454, 670,
  454, 670, 474, 650, 478, 646, 454, 674,
  450, 674, 450, 1798, 474, 650, 450, 1798,
  454, 1794, 454, 1794, 454, 1798, 450, 1798,
  450, 1798, 474, 1000
};

void sent() {
  Serial.println("sent");
  IrSender.sendSamsung(uint16_t aAddress, uint16_t aCommand, int_fast8_t aNumberOfRepeats)
    IrSender.sendRaw(rawData, sizeof(rawData) / sizeof(rawData[0]), 36);  // Note the approach used to automatically calculate the size of the array.

  //
  // receiver.enableIRIn();
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
