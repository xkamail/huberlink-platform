package thing

import (
	"context"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var ErrConnection = uierr.Alert("connection to device error")

const (
	PrefixTopic = "huberlink"
)

var cfg = config.Load()

func New() (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.MQTTURI)
	opts.SetUsername(cfg.MQTTUsername)
	opts.SetPassword(cfg.MQTTPassword)
	opts.SetClientID("")
	opts.SetCleanSession(true)
	opts.SetKeepAlive(60 * time.Second)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}

type Subscriber interface {
	// Topic return topic name to subscribe
	Topic() string
	// Handler return handler function
	Handler(ctx context.Context, deviceID snowid.ID, p []byte) error
}

type ReportFormat[T any] struct {
	Time int64 `json:"time"`
	Data T     `json:"data"`
}

func Call(ctx context.Context, topic string, deviceID snowid.ID, payload []byte) error {
	c, err := New()
	if err != nil {
		return err
	}
	defer c.Disconnect(250)
	topic = fmt.Sprintf("%s/%s/%s", PrefixTopic, deviceID.String(), topic)
	token := c.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}
