#define F_CPU 8000000UL
#ifndef __AVR_ATmega328P__
#define __AVR_ATmega328P__
#endif
#include <avr/io.h>
#include <avr/interrupt.h>
#include <string.h> //needed for memcpy
#include <SoftwareSerial.h>

SoftwareSerial _esp(PIN6, PIN7); // RX, TX

// declare functions
uint16_t readAnalog(uint8_t p = 0);

void setup()
{
  while (!Serial)
  {
  }
  Serial.begin(9600);
  _esp.begin(9600);
  //
  PORTD = 0x00;
  DDRD = 0x00;

  initTimer();
  // ok lets interrupt
  sei();
  sendCMD("AT+RST");
  // delay(1000);
}

void loop()
{
  if (Serial.available())
  {
    delay(1000);
    String command = "";
    while (Serial.available())
    {
      command += (char)Serial.read();
    }
    Serial.println(command);
    _esp.println(command);
  }
  if (_esp.available() > 0)
  {
    //
    while (_esp.available())
    {
      Serial.print("esp: ");
      char c = _esp.read();
      Serial.write(c);
    }
  }
}

void initTimer()
{
  // set timer2 interrupt at 1Hz
  TCCR2A = 0; // set entire TCCR2A register to 0
  TCCR2B = 0; // same for TCCR2B
  TCNT2 = 0;  // initialize counter value to 0
  // set compare match register for 1hz increments
  OCR2A = 249; // = (16*10^6) / (1000*64) - 1 (must be <256)
  // turn on CTC mode
  TCCR2A |= (1 << WGM21);
  // set clock source to 1024 prescaler
  TCCR2B |= (1 << CS22) | (1 << CS21) | (1 << CS20);
  // enable timer compare interrupt
  // TIMSK2 |= (1 << OCIE2A);
  return;
}

ISR(TIMER2_COMPA_vect)
{
  // int x = analogRead(PIN_A0);
  // Serial.println(x);
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

bool sendCMD(String cmd)
{
  Serial.println("sendCMD: " + cmd);
  _esp.println(cmd);
  return true;
}
