package device

import (
	"context"
	"time"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/iot/home"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/rand"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

type Device struct {
	ID                snowid.ID  `json:"id"`
	Name              string     `json:"name"`
	Icon              string     `json:"icon"`
	Model             string     `json:"model"`
	Kind              Kind       `json:"kind"`
	HomeID            snowid.ID  `json:"homeId"`
	UserID            snowid.ID  `json:"userId"`
	Token             string     `json:"token"`
	IpAddress         *string    `json:"ipAddress"`
	Location          *string    `json:"location"`
	LatestHeartbeatAt *time.Time `json:"latestHeartbeatAt"`
	CreatedAt         string     `json:"createdAt"`
	UpdatedAt         string     `json:"updatedAt"`
}

type CreateParam struct {
	Name   string    `json:"name"`
	Icon   string    `json:"icon"`
	Model  string    `json:"model"`
	Kind   Kind      `json:"kind"`
	HomeID snowid.ID `json:"-"`
}

func Create(ctx context.Context, p *CreateParam) (*snowid.ID, error) {
	acc, err := account.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	h, err := home.GetFromIDAndUserID(ctx, p.HomeID, acc.ID)
	if err != nil {
		return nil, err
	}
	deviceToken, err := generateToken()
	if err != nil {
		return nil, err
	}

	// open transaction
	// insert device
	tx, err := pgctx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var id snowid.ID
	err = tx.QueryRow(ctx, `
		insert into devices
			(id, name, icon, model, kind, home_id, user_id, token, ip_address, location, latest_heartbeat_at, created_at, updated_at) 
		values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $12) 
		returning id`,
		snowid.Gen(),
		p.Name,
		p.Icon,
		p.Model,
		p.Kind,
		h.ID,
		acc.ID,
		deviceToken,
		nil,
		nil,
		nil,
		time.Now(),
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	switch p.Kind {
	case KindIRRemote:
		_, err = tx.Exec(ctx, `
			insert into device_ir_remotes 
				(id, device_id, home_id, created_at, updated_at) 
			values ($1,$2,$3,$4,$4)`,
			snowid.Gen(), // id
			id,           // device_id
			h.ID,         // home_id
			time.Now(),   // created_at, updated_at
		)
		if err != nil {
			return nil, err
		}
	case KindLamp:

	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &id, nil
}

func generateToken() (string, error) {
	s, err := rand.String(500)
	if err != nil {
		return "", err
	}
	return s, nil
}
