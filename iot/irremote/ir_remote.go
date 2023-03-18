package irremote

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

// VirtualCategory is an enum that represents a virtual key category
type VirtualCategory uint

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
	Icon     string          `json:"icon"`
	Kind     VirtualCategory `json:"category"`
	// IsLearning is a flag that indicates that the virtual remote is learning
	// when rawData codes has come
	// Command will be created and IsLearning will be false
	//
	IsLearning bool      `json:"isLearning"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Command struct {
	ID        snowid.ID `json:"id"`
	RemoteID  snowid.ID `json:"remoteId"`
	VirtualID snowid.ID `json:"virtualId"`
	Name      string    `json:"name"`
	Code      []uint    `json:"code"`
	Remark    *string   `json:"remark"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func Find(ctx context.Context, deviceID snowid.ID) (*IRRemote, error) {
	// TODO: implement
	panic("not implemented")
}

func ListVirtual(ctx context.Context, deviceID snowid.ID) ([]*VirtualKey, error) {
	// TODO: implement
	panic("not implemented")
}

type CreateVirtualKeyParam struct {
	Name     string `json:"name"`
	Kind     string `json:"kind"`
	Icon     string `json:"icon"`
	remoteID snowid.ID
}

func CreateVirtual(ctx context.Context, p *CreateVirtualKeyParam) (*VirtualKey, error) {
	if _, err := account.FromContext(ctx); err != nil {
		return nil, err
	}

	remoteID := p.remoteID

	err := pgctx.QueryRow(ctx, `select id from device_ir_remote_virtual_keys where remote_id = $1`,
		remoteID, //
	).Scan(&remoteID)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	tx, err := pgctx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	id := snowid.Gen()
	_, err = tx.Exec(ctx, `
			insert into device_ir_remote_virtual_keys 
			(id, remote_id, name, kind, icon, created_at, updated_at) 
			values ($1, $2, $3, $4, $5, $6, $7)`,
		id,
		remoteID,
		p.Name,
		p.Kind,
		p.Icon,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return FindVirtual(ctx, remoteID, id)
}

func FindVirtual(ctx context.Context, remoteID, virtualID snowid.ID) (*VirtualKey, error) {
	// TODO: test this function
	rows, err := pgctx.Query(ctx, `select id, remote_id, name, kind, icon, created_at, updated_at from device_ir_remote_virtual_keys where remote_id = $1 and id = $2`, remoteID, virtualID)
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
