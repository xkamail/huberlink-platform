package irremote

import (
	"context"
	"encoding/json"

	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/thing"
)

type MQTTCmd struct {
	Code      string `json:"code"`
	Frequency int    `json:"frequency"`
}

type MQTTReport struct {
	VirtualKeyID snowid.ID  `json:"virtualKeyId"`
	Properties   Properties `json:"properties"`
}

func NewSubscribeLearning() thing.Subscriber {
	return &learningSubscribe{}
}

var _ thing.Subscriber = (*learningSubscribe)(nil)

type learningSubscribe struct {
}

func (t learningSubscribe) Topic() string {
	return "irremote/learning"
}

type MQTTLearningPayload struct {
	DeviceID snowid.ID `json:"deviceId"`
	RawData  []uint8   `json:"rawData"`
	Platform string    `json:"platform"`
}

func (t learningSubscribe) Handler(ctx context.Context, _b []byte) error {
	var p MQTTLearningPayload
	if err := json.Unmarshal(_b, &p); err != nil {
		return err
	}
	// find  virtual key which is learning state
	command, err := CreateCommand(ctx, &CreateCommandParam{
		Name:     "",
		Remark:   "",
		DeviceID: p.DeviceID,
		Platform: p.Platform,
		Code:     p.RawData,
	})
	if err != nil {
		slog.Debug("create command error", err)
		return err
	}
	slog.Debug("create command", slog.Any("command", command))
	return nil
}
