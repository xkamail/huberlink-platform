#include <Arduino.h>
#include <SPI.h>
#include "avr/interrupt.h"
#define IR_SEND_PIN 3

#include <IRremote.hpp>
/*
 * Helper macro for getting a macro definition as string
 */
#if !defined(STR_HELPER)
#define STR_HELPER(x) #x
#define STR(x) STR_HELPER(x)
#endif
void setup() {
  IrSender.begin(3);
  pinMode(SS, INPUT);
  while (!Serial) {}
  Serial.begin(9600);
  Serial.println("Start");
  Serial.print(F("Ready to send IR signals at pin " STR(IR_SEND_PIN) " on press of button at pin "));
  SPCR |= _BV(SPE);   /* Enable SPI */
  DDRB = (1 << PB4);  // slave out
  SPI.attachInterrupt();
}

uint8_t irData[600];  // max buffer remote code
volatile bool _success = false;
volatile int _length = 0;  // true data length

void loop() {
  if (_success) {
    _success = false;
    int m = _length;
    _length = 0;
    sent(m);
  }
}

void sent(int _length) {
  IrSender.sendRaw(irData, _length, 38);
  Serial.println("[INFO] sent");
}

ISR(SPI_STC_vect) {
  uint8_t oldsrg = SREG;
  cli();
  uint8_t n = (uint8_t)SPDR;
  if (n == 0) {
    _success = true;
    SREG = oldsrg;
    return;
  }
  irData[_length++] = n;
  SREG = oldsrg;
}
