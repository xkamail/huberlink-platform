package irremote

import (
	"context"
	"strconv"
	"strings"

	"golang.org/x/exp/slog"

	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/thing"
)

const ExecuteTopic = "thing/irremote/execute"
const LearningTopic = "thing/irremote/learning"

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
	return "irremote/learning/result"
}

type MQTTLearningPayload struct {
	RawData  []uint8 `json:"rawData"`
	Platform string  `json:"platform"`
}

func (t learningSubscribe) Handler(ctx context.Context, deviceID snowid.ID, _b []byte) error {
	codes := make([]uint8, 0)
	str := string(_b)
	slog.Debug("learning result", slog.String("str", str))
	for _, c := range strings.Split(str, ",") {
		if c == "" {
			continue
		}
		atoi, err := strconv.Atoi(c)
		if err != nil {
			slog.Error("atoi error", err)
			return err
		}
		codes = append(codes, uint8(atoi))
	}

	// find  virtual key which is learning state
	command, err := CreateCommand(ctx, &CreateCommandParam{
		Name:     "", // user have to setting name after learning
		Remark:   "", // for specific frontend to show
		DeviceID: deviceID,
		Platform: "",
		Code:     codes,
	})
	if err != nil {
		slog.Debug("create command error", err)
		return err
	}
	slog.Debug("create command", slog.Any("command", command))
	return nil
}
