package irremote

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

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
	// unknown, samsung, lg, ...
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type CreateCommandParam struct {
	Name   string
	Remark string
	Code   []uint8
	// from url params
	Platform string
	DeviceID snowid.ID
}

// CreateCommand will be trigger when thing. topic is learning
func CreateCommand(ctx context.Context, p *CreateCommandParam) (*Command, error) {
	var (
		virtualID snowid.ID
		remoteID  snowid.ID
	)
	// find  virtual key which is learning state
	err := pgctx.QueryRow(ctx, `
			select is_learning, dir.id 
			from device_ir_remote_virtual_keys 
			inner join device_ir_remotes dir on dir.id = device_ir_remote_virtual_keys.remote_id
			where dir.device_id = $1 and is_learning = true`,
		p.DeviceID,
	).Scan(
		&virtualID,
		&remoteID,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, uierr.NotFound("no virtual key is learning")
	}
	if err != nil {
		return nil, err
	}
	tx, err := pgctx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var cmdID snowid.ID

	err = tx.QueryRow(ctx, `
		insert into device_ir_remote_commands 
			(id, remote_id, virtual_id, name, code, remark, platforms, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
		returning id`,
		snowid.Gen(),
		remoteID,
		virtualID,
		p.Name,
		p.Code, // wait for is_learning mode
		p.Remark,
		p.Platform,
		time.Now(),
		time.Now(),
	).Scan(&cmdID)
	if err != nil {
		return nil, err
	}
	// commit transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return nil, nil
}

func FindCommand(ctx context.Context, deviceID, virtualID, commandID snowid.ID) (*Command, error) {
	return nil, nil
}

func ListCommand(ctx context.Context, deviceID, virtualID snowid.ID) ([]*Command, error) {
	return nil, nil
}

func DeleteCommand(ctx context.Context, deviceID, virtualID, commandID snowid.ID) error {
	return nil
}
