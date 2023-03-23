#include <IRremote.hpp>
#define MARK_EXCESS_MICROS 20  // Adapt it to your IR receiver module. 20 is recommended for the cheap VS1838 modules.

#define RAW_BUFFER_LENGTH 400

struct storedIRDataStruct {
  IRData receivedIRData;
  // extensions for sendRaw
  uint8_t rawCode[RAW_BUFFER_LENGTH];  // The durations if raw
  uint8_t rawCodeLength;               // The length of the code
} sStoredIRData;

void setup() {
  //
  Serial.begin(9600);
  attachInterrupt(1, sw, RISING);
  IrSender.begin(4);
  IrReceiver.begin(PIND2, true);
  Serial.println("begin");
}

void loop() {

  //
  if (IrReceiver.decode()) {
    storeCode();
    IrReceiver.resume();  // resume receiver
  }
}

void sw() {
  IrReceiver.stop();
  Serial.println("click");
  sendCode(&sStoredIRData);
}

// Stores the code for later playback in sStoredIRData
// Most of this code is just logging
void storeCode() {
  Serial.println("storeCode");
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
  /*
     * Copy decoded data
     */
  sStoredIRData.receivedIRData = IrReceiver.decodedIRData;

  if (sStoredIRData.receivedIRData.protocol == UNKNOWN) {
    Serial.print(F("Received unknown code and store "));
    Serial.print(IrReceiver.decodedIRData.rawDataPtr->rawlen - 1);
    Serial.println(F(" timing entries as raw "));
    IrReceiver.printIRResultRawFormatted(&Serial, true);  // Output the results in RAW format
    sStoredIRData.rawCodeLength = IrReceiver.decodedIRData.rawDataPtr->rawlen - 1;
    /*
         * Store the current raw data in a dedicated array for later usage
         */
    IrReceiver.compensateAndStoreIRResultInArray(sStoredIRData.rawCode);
  } else {
    sStoredIRData.rawCodeLength = IrReceiver.decodedIRData.rawDataPtr->rawlen - 1;
    IrReceiver.compensateAndStoreIRResultInArray(sStoredIRData.rawCode);
    Serial.println("got known value");
    sStoredIRData.receivedIRData.flags = 0;  // clear flags -esp. repeat- for later sending
    Serial.println();
  }
}


void sendCode(storedIRDataStruct *aIRDataToSend) {
  Serial.println("sent");

  IrSender.sendRaw(aIRDataToSend->rawCode, aIRDataToSend->rawCodeLength, 38);
  Serial.println("===");
  for (int i = 0; i < aIRDataToSend->rawCodeLength; i++) {
    //
    Serial.print(aIRDataToSend->rawCode[i]);
    if (i != aIRDataToSend->rawCodeLength - 1) {
      Serial.print(",");
    }
  }
  Serial.println("");
}