package home

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/iot/device"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

type SceneExecuter interface {
	ExecuteScene(scene Scene) error
}
type SceneRun uint

const (
	SceneRunWhenTap SceneRun = 1 << iota
	SceneRunWhenSchedule
)

type SceneScheduleRepeat uint

const (
	SceneScheduleRepeatSun = 1 << iota
	SceneScheduleRepeatMon
	SceneScheduleRepeatTue
	SceneScheduleRepeatWed
	SceneScheduleRepeatThu
	SceneScheduleRepeatFri
	SceneScheduleRepeatSat
)

type SceneScheduleTime string

func (t *SceneScheduleTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(*t))
}

func (t *SceneScheduleTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	// 00:00 is valid
	if s == "00:00" {
		*t = SceneScheduleTime(s)
		return nil
	}
	// hh:mm
	if len(s) != 5 {
		return uierr.BadInput("time", "invalid time format")
	}
	if s[2] != ':' {
		return uierr.BadInput("time", "invalid time format")
	}
	hh, err := strconv.ParseInt(s[:2], 10, 64)
	if err != nil {
		return uierr.BadInput("time", "invalid time format: hh should be 0-23")
	}
	// validate hh should be 0-23
	mm, err := strconv.ParseInt(s[3:], 10, 64)
	if err != nil {
		return uierr.BadInput("time", "invalid time format: mm should be 0-59")
	}
	// validate mm should be 0-59
	if hh < 0 || hh > 23 || mm < 0 || mm > 59 {
		return uierr.BadInput("time", "invalid time format: hh:mm should be 0-23:0-59")
	}

	*t = SceneScheduleTime(s)
	return nil
}

var _ json.Marshaler = (*SceneScheduleTime)(nil)
var _ json.Unmarshaler = (*SceneScheduleTime)(nil)

type Scene struct {
	ID     snowid.ID `json:"id"`
	HomeID snowid.ID `json:"homeId"`
	Name   string    `json:"name"`
	Run    SceneRun  `json:"run"`
	// when no repeat is zero value
	// multiple days can be set by bitwise or
	ScheduleRepeat SceneScheduleRepeat `json:"schedule"`
	// hh:mm format
	ScheduleTime SceneScheduleTime `json:"scheduleTime"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SceneAction struct {
	ID       snowid.ID `json:"id"`
	Order    uint      `json:"order"`
	Name     string    `json:"name"`
	DeviceID snowid.ID `json:"deviceId"` // device id to execute
	// action to execute
	RawAction string `json:"rawAction"`
}

type CreateSceneParam struct {
	Name string   `json:"name"`
	Run  SceneRun `json:"run"`
	// when no repeat is zero value
	// multiple days can be set by bitwise or
	ScheduleRepeat SceneScheduleRepeat `json:"schedule"`
	// hh:mm format
	ScheduleTime SceneScheduleTime `json:"scheduleTime"`
}

func CreateScene(ctx context.Context, homeID snowid.ID, p *CreateSceneParam) (snowid.ID, error) {
	panic("implement me")
}

func ListScene(ctx context.Context, homeID snowid.ID) ([]*Scene, error) {
	panic("implement me")
}

type UpdateSceneParam = CreateSceneParam

func UpdateScene(ctx context.Context, homeID snowid.ID, p *UpdateSceneParam) error {
	panic("implement me")
}

func DeleteScene(ctx context.Context, homeID, sceneID snowid.ID) error {
	panic("implement me")
}

type CreateSceneActionParam struct {
	SceneID snowid.ID `json:"sceneId"`
	// DeviceID in home id
	DeviceID  snowid.ID `json:"deviceId"`
	RawAction string    `json:"rawAction"`
}

func CreateSceneAction(ctx context.Context, homeID snowid.ID, p *CreateSceneActionParam) (snowid.ID, error) {
	err := pgctx.QueryRow(ctx, `select id from devices where id = $1 and home_id = $2`,
		p.DeviceID,
		homeID,
	).Scan(
		&p.DeviceID,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, device.ErrNotFound
	}
	if err != nil {
		return 0, err
	}
	var id snowid.ID
	err = pgctx.QueryRow(ctx, `
		insert into home_scenes_actions 
		    (id, scene_id, device_id, action) 
		values ($1,$2,$3) returning id`,
		snowid.New(),
		homeID,
		p.SceneID,
		p.DeviceID,
		p.RawAction,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteSceneAction(ctx context.Context, homeID, sceneActionID snowid.ID) error {
	panic("implement me")
}
