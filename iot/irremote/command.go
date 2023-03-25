package irremote

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/thing"
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
			select device_ir_remote_virtual_keys.id, dir.id 
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
	_, err = tx.Exec(ctx, `update device_ir_remote_virtual_keys set is_learning = false where id = $1`, virtualID)
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
	rows, err := pgctx.Query(ctx, `
		select c.id, c.remote_id, c.virtual_id, c.name, c.code, c.remark, c.platforms, c.created_at, c.updated_at 
		from device_ir_remote_commands c 
		inner join device_ir_remotes dir on dir.id = c.remote_id 
		where c.id = $1 and c.virtual_id = $2 and dir.device_id = $3`,
		commandID,
		virtualID,
		deviceID,
	)
	if err != nil {
		return nil, err
	}
	c, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[Command])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrCommandNotFound
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ListCommand(ctx context.Context, deviceID, virtualID snowid.ID) ([]*Command, error) {
	rows, err := pgctx.Query(ctx, `
		select c.id, c.remote_id, c.virtual_id, c.name, c.code, c.remark, c.platforms, c.created_at, c.updated_at 
		from device_ir_remote_commands c 
		inner join device_ir_remotes dir on dir.id = c.remote_id 
		where c.virtual_id = $2 and dir.device_id = $3`,
		virtualID,
		deviceID,
	)
	if err != nil {
		return nil, err
	}
	c, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByPos[Command])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrCommandNotFound
	}
	if err != nil {
		return nil, err
	}
	return c, nil
}

func DeleteCommand(ctx context.Context, deviceID, virtualID, commandID snowid.ID) error {
	_, err := pgctx.Exec(ctx, `
		delete from device_ir_remote_commands
		where id = $1 and virtual_id = $2`,
		commandID,
		virtualID,
	)
	return err
}

type UpdateCommandParam struct {
	Name   string `json:"name"`
	Remark string `json:"remark"`
}

func (p *UpdateCommandParam) Valid() error {
	if p.Name == "" {
		return uierr.Invalid("name", "name is required")
	}

	return nil
}

func UpdateCommand(ctx context.Context, deviceID, virtualID, commandID snowid.ID, p *UpdateCommandParam) (*Command, error) {

	_, err := pgctx.Exec(ctx, `
		update device_ir_remote_commands 
			set name = $1, 
			remark = $2 
			where id = $3 and virtual_id = $4 `,
		p.Name,
		p.Remark,
		commandID,
		virtualID,
	)
	if err != nil {
		return nil, err
	}
	return FindCommand(ctx, deviceID, virtualID, commandID)
}

type ExecuteCommandParam struct {
	CommandID snowid.ID `json:"commandId"`
}
type ExecuteResult struct {
	//
}

func ExecuteCommand(ctx context.Context, deviceID, virtualID snowid.ID, p *ExecuteCommandParam) (*ExecuteResult, error) {
	var codes []uint
	err := pgctx.QueryRow(ctx, `
		select code 
		from device_ir_remote_commands inner join device_ir_remotes dir on dir.id = device_ir_remote_commands.remote_id 
		where virtual_id = $1 and dir.device_id = $2 and device_ir_remote_commands.id = $3`,
		virtualID,
		deviceID,
		p.CommandID,
	).Scan(&codes)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrCommandNotFound
	}
	if err != nil {
		return nil, err
	}

	codeString := make([]string, 0)
	// convert array of uint to string
	for i := 0; i < len(codes); i++ {
		codeString = append(codeString, strconv.Itoa(int(codes[i])))
	}
	payload := strings.Join(codeString, ",")

	if err := thing.Call(ctx, ExecuteTopic, deviceID, []byte(payload)); err != nil {
		return nil, err
	}

	return &ExecuteResult{
		// TODO
	}, nil
}
