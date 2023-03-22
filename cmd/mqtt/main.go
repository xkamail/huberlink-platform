package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/iot/device"
	"github.com/xkamail/huberlink-platform/iot/irremote"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/thing"
)

var debug = flag.Bool("debug", false, "enable debug log")

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run MQTT server
func run() error {
	l := slog.NewTextHandler(os.Stdout)
	if *debug {
		l.Enabled(context.TODO(), slog.LevelDebug)
	}
	slog.SetDefault(slog.New(l))

	ctx := context.Background()
	if err := config.Init(); err != nil {
		return err
	}
	cfg := config.Load()

	conn, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer conn.Close()
	if err := conn.Ping(context.Background()); err != nil {
		return err
	}
	// db session to context
	ctx = pgctx.NewContext(ctx, conn)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(cfg.MQTTURI)
	opts.SetUsername(cfg.MQTTUsername)
	opts.SetPassword(cfg.MQTTPassword)
	opts.SetClientID("")
	opts.SetCleanSession(true)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetDefaultPublishHandler(handler(ctx))
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	if err != nil {
		return err
	}
	err = thing.Execute(context.Background(), snowid.Gen(), &thing.ExecuteMessage{
		Time: time.Now().Unix(),
		Data: &irremote.MQTTCmd{
			Code:      "0x123",
			Frequency: 38,
		},
	})
	if err != nil {
		return err
	}
	c.IsConnected()
	log.Printf("server is running %v\n", c.IsConnected())

	select {}
}

/*
 tylink/${deviceId}/thing/property/report
 tylink/${deviceId}/thing/property/set
 tylink/${deviceId}/thing/property/get
*/

// handler is a default handler for all messages
func handler(ctx context.Context) func(client mqtt.Client, msg mqtt.Message) {
	irRemoteLearningSub := irremote.NewSubscribeLearning()
	fn := func(client mqtt.Client, msg mqtt.Message) error {
		deviceID, topic, err := ExtractTopic(msg.Topic())
		if err != nil {
			return err
		}
		iot, err := device.Find(ctx, deviceID)
		if err != nil {
			return err
		}
		var payload thing.ReportFormat[any]
		if err := json.Unmarshal(msg.Payload(), &payload); err != nil {
			return err
		}

		switch iot.Kind {
		// we only support IRRemote for now
		case device.KindIRRemote:
			if irRemoteLearningSub.Topic() == topic {
				return irRemoteLearningSub.Handler(ctx, msg.Payload())
			}
		default:
			return fmt.Errorf("not support device kind %d", iot.Kind)
		}
		return nil
	}
	return func(client mqtt.Client, msg mqtt.Message) {
		slog.Debug("Received message on topic: %s\nMessage: %s\n", msg.Topic(), string(msg.Payload()))
		//
		err := fn(client, msg)
		if err != nil {
			slog.Error("handler error", err)
			return
		}
		// acknowledge message when no error occurred
		msg.Ack()
	}
}

// ExtractTopic extract device id and real_topic from topic
func ExtractTopic(_topic string) (snowid.ID, string, error) {
	deviceID, err := deviceIDFromTopic(_topic)
	if err != nil {
		return 0, "", err
	}

	delimiter := fmt.Sprintf("%s/%s/thing/", thing.PrefixTopic, deviceID.String())
	s := strings.TrimPrefix(_topic, delimiter)
	if len(s) == 0 {
		return 0, "", fmt.Errorf("invalid topic: no real topic found %s", _topic)
	}
	return deviceID, s, nil
}

func deviceIDFromTopic(topic string) (snowid.ID, error) {
	if !strings.HasPrefix(topic, fmt.Sprintf("%s/", thing.PrefixTopic)) {
		return snowid.Zero, fmt.Errorf("invalid topic %s", topic)
	}
	s := strings.Split(topic, "/")
	if len(s) < 2 {
		return snowid.Zero, fmt.Errorf("invalid topic len %s", topic)
	}
	return snowid.Parse(s[1])
}
