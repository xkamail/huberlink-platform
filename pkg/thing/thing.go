package thing

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

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
	Handler(ctx context.Context, p []byte) error
}

type ReportFormat[T any] struct {
	Time int64 `json:"time"`
	Data T     `json:"data"`
}

type ExecuteMessage struct {
	Time int64 `json:"time"`
	Data any   `json:"data"`
}

func Execute(ctx context.Context, deviceId snowid.ID, m *ExecuteMessage) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		b, err := json.Marshal(m)
		if err != nil {
			return err
		}
		c, err := New()
		if err != nil {
			return err
		}
		defer c.Disconnect(250)
		topic := fmt.Sprintf("%s/%s/thing/execute", PrefixTopic, deviceId.String())

		c.Publish(topic, 1, false, string(b)).
			WaitTimeout(2 * time.Second)
		return nil
	}
}
