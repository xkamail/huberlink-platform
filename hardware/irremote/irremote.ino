#define F_CPU 8000000UL
#ifndef __AVR_ATmega328P__
#define __AVR_ATmega328P__
#endif
#include <avr/io.h>
#include <avr/interrupt.h>
#include <string.h> //needed for memcpy
#include <SoftwareSerial.h>

// declare functions
uint16_t readAnalog(uint8_t p = 0);

void setup()
{
  while (!Serial)
  {
  }
  Serial.begin(9600);
  //
  PORTD = 0x00;
  DDRD = 0x00;

  initTimer();
  // ok lets interrupt
  sei();
}

void loop()
{
  uint16_t x = readAnalog(0);
  Serial.println();
  Serial.println(x);
  Serial.println(analogRead(A0));
  Serial.println();
  delay(1000);
}

ISR(TIMER2_COMPA_vect)
{
  //
}

uint16_t readAnalog(uint8_t p = 0)
{
  // select pin
  ADMUX = (1 << REFS0) | (p & 0x07);
  ADCSRA |= (1 << ADEN);
  // start conversion
  ADCSRA |= (1 << ADSC);
  while (ADCSRA & (1 << ADSC))
  {
  }
  // re allocate uint16_t

  uint16_t result = (ADCH << 8) | (ADCL);
  return result;
}
