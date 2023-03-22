package irremote

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

type Virtualer interface {
	// Properties get default properties
	Properties() Properties
	// PropertiesKeys validate properties
	PropertiesKeys() []string
}

// VirtualCategory is an enum that represents a virtual key category
type VirtualCategory uint

func (v *VirtualCategory) UnmarshalJSON(b []byte) error {
	u, err := strconv.ParseUint(string(b), 10, 8)
	if err != nil {
		return err
	}
	if u > uint64(VirtualCategoryWaterHeater) {
		return uierr.Invalid("category", "invalid virtual category")
	}
	*v = VirtualCategory(u)
	return nil
}

func (v *VirtualCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(*v)
}

var _ json.Marshaler = (*VirtualCategory)(nil)
var _ json.Unmarshaler = (*VirtualCategory)(nil)

// Properties represent state of a virtual key (device)
type Properties map[string]string

const (
	VirtualCategoryOther VirtualCategory = iota
	VirtualCategoryTV
	VirtualCategoryAirConditioner
	VirtualCategoryLight
	VirtualCategoryFan
	VirtualCategorySpeaker
	VirtualCategoryProjector
	VirtualCategoryDVD
	VirtualCategoryWaterHeater
)

// IRRemote is a struct that represents a remote control
// that linked to a device.Device and home.Home
type IRRemote struct {
	ID        snowid.ID `json:"id"`
	DeviceID  snowid.ID `json:"deviceId"`
	HomeID    snowid.ID `json:"homeId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// VirtualKey is a struct that represents a remote control
// virtual remote control
type VirtualKey struct {
	ID       snowid.ID       `json:"id"`
	RemoteID snowid.ID       `json:"remoteId"`
	Name     string          `json:"name"`
	Kind     VirtualCategory `json:"category"`
	Icon     string          `json:"icon"`
	// IsLearning is a flag that indicates that the virtual remote is learning
	// when rawData codes has come
	// Command will be created and IsLearning will be false
	//
	IsLearning bool       `json:"isLearning"`
	Properties Properties `json:"properties"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

type Command struct {
	ID        snowid.ID `json:"id"`
	RemoteID  snowid.ID `json:"remoteId"`
	VirtualID snowid.ID `json:"virtualId"`
	Name      string    `json:"name"`
	// Code of raw data
	Code []uint `json:"-"`
	// remark is a note for frontend
	Remark *string `json:"remark"`
	// Platform is a platform that this command will be used
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Find IRRemote by deviceID
func Find(ctx context.Context, deviceID snowid.ID) (*IRRemote, error) {
	rows, err := pgctx.Query(ctx, `select id, device_id, home_id, created_at, updated_at from device_ir_remotes where device_id = $1`, deviceID)
	if err != nil {
		return nil, err
	}
	d, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[IRRemote])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

func ListVirtual(ctx context.Context, deviceID snowid.ID) ([]*VirtualKey, error) {
	rows, err := pgctx.Query(ctx, `select v.id, v.remote_id, v.name, v.kind, v.icon, v.is_learning, v.properties, v.created_at, v.updated_at 
		from device_ir_remote_virtual_keys v inner join device_ir_remotes dir on dir.id = v.remote_id 
		where dir.device_id = $1`,
		deviceID,
	)
	if err != nil {
		return nil, err
	}
	d, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[VirtualKey])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrVirtualKeyNotfound
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

type CreateVirtualKeyParam struct {
	Name     string          `json:"name"`
	Kind     VirtualCategory `json:"kind"`
	Icon     string          `json:"icon"`
	DeviceID snowid.ID       `json:"-"`
}

func CreateVirtual(ctx context.Context, p *CreateVirtualKeyParam) (*VirtualKey, error) {
	if _, err := account.FromContext(ctx); err != nil {
		return nil, err
	}
	remoteID, err := findRemoteID(ctx, p.DeviceID)
	if err != nil {
		return nil, err
	}
	tx, err := pgctx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	properties := make(Properties)

	switch p.Kind {
	case VirtualCategoryTV:
		properties = TV{}.Properties()
	case VirtualCategoryAirConditioner:
		properties = Air{}.Properties()
	case VirtualCategoryLight:
	case VirtualCategoryFan:
	case VirtualCategorySpeaker:
	case VirtualCategoryProjector:
	case VirtualCategoryDVD:
	case VirtualCategoryWaterHeater:
	case VirtualCategoryOther:
		properties = Other{}.Properties()
	default:
	}

	id := snowid.Gen()
	_, err = tx.Exec(ctx, `
			insert into device_ir_remote_virtual_keys 
			(id, remote_id, name, kind, icon, is_learning, properties, created_at, updated_at) 
			values ($1, $2, $3, $4, $5, false, $6, $7, $8)`,
		id,
		remoteID,
		p.Name,
		p.Kind,
		p.Icon,
		properties,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return FindVirtual(ctx, p.DeviceID, id)
}

type UpdateVirtualParam struct {
	Name      string    `json:"name"`
	VirtualID snowid.ID `json:"-"`
}

func UpdateVirtual(ctx context.Context, p *UpdateVirtualParam) error {
	if _, err := account.FromContext(ctx); err != nil {
		return err
	}

	tx, err := pgctx.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `
			update device_ir_remote_virtual_keys 
			set name = $1, updated_at = $2
			where id = $3`,
		p.Name,
		time.Now(),
		p.VirtualID,
	)
	if err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func FindVirtual(ctx context.Context, deviceID, virtualID snowid.ID) (*VirtualKey, error) {
	// TODO: test this function
	rows, err := pgctx.Query(ctx, `select v.id, remote_id, name, kind, icon, is_learning, properties, v.created_at, v.updated_at from device_ir_remote_virtual_keys v inner join device_ir_remotes dir on dir.id = v.remote_id 
		where dir.device_id = $1 and v.id = $2`,
		deviceID,
		virtualID,
	)
	if err != nil {
		return nil, err
	}
	d, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[VirtualKey])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrVirtualKeyNotfound
	}
	if err != nil {
		return nil, err
	}

	return d, nil
}

func DeleteVirtualKey(ctx context.Context, virtualID snowid.ID) (bool, error) {
	_, err := pgctx.Exec(ctx, `delete from device_ir_remote_virtual_keys where id = $1`, virtualID)
	return err == nil, err
}

func ListVirtualCommand(ctx context.Context, virtualID snowid.ID) ([]*Command, error) {
	rows, err := pgctx.Query(ctx, `select id, remote_id, virtual_id, name, code, remark, created_at, updated_at from device_ir_remote_commands where virtual_id = $1`, virtualID)
	if err != nil {
		return nil, err
	}
	d, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Command])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrCommandNotfound
	}
	if err != nil {
		return nil, err
	}
	return d, nil
}

func findRemoteID(ctx context.Context, deviceID snowid.ID) (snowid.ID, error) {
	var remoteID snowid.ID
	err := pgctx.QueryRow(ctx, `select id from device_ir_remotes where device_id = $1`,
		deviceID, //
	).Scan(&remoteID)
	if errors.Is(err, pgx.ErrNoRows) {
		return snowid.Zero, ErrNotFound
	}
	if err != nil {
		return snowid.Zero, err
	}
	return remoteID, nil
}
