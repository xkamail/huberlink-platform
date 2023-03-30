package home

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrSceneActionNotFound = uierr.NotFound("scene action not found")
	ErrSceneNotFound       = uierr.NotFound("scene not found")
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

func FindScene(ctx context.Context, homeID, sceneID snowid.ID) (*Scene, error) {
	rows, err := pgctx.Query(ctx, `select * from home_scenes where home_id = $1 and id = $2`,
		homeID,
		sceneID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	s, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[Scene])
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrSceneNotFound
	}
	if err != nil {
		return nil, err
	}
	return s, nil
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
	var id snowid.ID
	err := pgctx.QueryRow(ctx, `
		insert into home_scenes 
		    (id, home_id, name, run, schedule_repeat, schedule_time) 
		values ($1,$2,$3,$4,$5,$6) returning id`,
		snowid.New(),
		homeID,
		p.Name,
		p.Run,
		p.ScheduleRepeat,
		p.ScheduleTime,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ListScene(ctx context.Context, homeID snowid.ID) ([]*Scene, error) {
	rows, err := pgctx.Query(ctx, `
		select id, name, run, schedule_repeat, schedule_time, created_at, updated_at
		from home_scenes
		where home_id = $1
		order by id desc`,
		homeID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var scenes []*Scene
	for rows.Next() {
		var s Scene
		err = rows.Scan(
			&s.ID,
			&s.Name,
			&s.Run,
			&s.ScheduleRepeat,
			&s.ScheduleTime,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		scenes = append(scenes, &s)
	}
	return scenes, nil
}

type UpdateSceneParam = CreateSceneParam

func UpdateScene(ctx context.Context, homeID, sceneID snowid.ID, p *UpdateSceneParam) error {
	_, err := pgctx.Exec(ctx, `
		update home_scenes set 
		    name = $1, run = $2, schedule_repeat = $3, schedule_time = $4
		where id = $5 and home_id = $6`,
		p.Name,
		p.Run,
		p.ScheduleRepeat,
		p.ScheduleTime,
		sceneID,
		homeID,
	)
	return err
}

func DeleteScene(ctx context.Context, homeID, sceneID snowid.ID) error {
	_, err := pgctx.Exec(ctx, `delete from home_scenes where id = $1 and home_id = $2`, sceneID, homeID)
	return err
}

type CreateSceneActionParam struct {
	DeviceID  snowid.ID `json:"deviceId"`
	RawAction string    `json:"rawAction"`
}

func CreateSceneAction(ctx context.Context, homeID, sceneID snowid.ID, p *CreateSceneActionParam) (snowid.ID, error) {
	err := pgctx.QueryRow(ctx, `select id from devices where id = $1 and home_id = $2`,
		p.DeviceID,
		homeID,
	).Scan(
		&p.DeviceID,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, uierr.NotFound("device not found")
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
		sceneID,
		p.DeviceID,
		p.RawAction,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func DeleteSceneAction(ctx context.Context, homeID, sceneID, sceneActionID snowid.ID) error {
	var c int
	err := pgctx.QueryRow(ctx, `select count(*) from home_scenes where home_id = $1 and id = $2`,
		homeID,
		sceneID,
	).Scan(
		&c,
	)
	if err != nil {
		return err
	}
	if c == 0 {
		return ErrSceneActionNotFound
	}
	_, err = pgctx.Exec(ctx, `delete from home_scenes_actions where id = $1 and scene_id = $2`,
		sceneActionID,
		sceneID,
	)
	return err
}

func ListSceneAction(ctx context.Context, homeID, sceneID snowid.ID) ([]*SceneAction, error) {
	rows, err := pgctx.Query(ctx, `
		select hsm.id, hsm.scene_id, hsm.device_id, hsm.action, hsm.created_at, hsm.updated_at, 
		from home_scenes_actions hsm 
		    inner join home_scenes hs on hs.id = hsm.scene_id 
		where scene_id = $1 and hs.home_id = $2 order by hsm.id desc`,
		sceneID,
		homeID,
	)
	if err != nil {
		return nil, err
	}
	actions, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[SceneAction])
	if err != nil {
		return nil, err
	}
	return actions, nil
}
