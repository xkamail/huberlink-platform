package thing

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

func New() (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://mc.xkamail.me:1883")
	opts.SetPassword("test")
	opts.SetUsername("test")
	opts.SetClientID("")
	opts.SetCleanSession(true)
	opts.SetKeepAlive(60 * time.Second)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
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
		topic := fmt.Sprintf("huberlink/%s/thing/property/report", deviceId.String())

		c.Publish(topic, 1, false, string(b)).
			WaitTimeout(2 * time.Second)
		return nil
	}
}
