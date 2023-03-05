package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// run MQTT server
func run() error {
	// create MQTT server
	opts := mqtt.NewClientOptions().
		AddBroker("tcp://mc.xkamail.me:1883").
		SetClientID("huberlink").
		SetPassword("huberlink").
		SetUsername("huberlink")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	select {}
}
