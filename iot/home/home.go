package home

import (
	"context"
	"strings"
	"time"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

type Home struct {
	ID            snowid.ID `json:"id"`
	Name          string    `json:"name"`
	UserID        snowid.ID `json:"userId"`
	BackgroundURL string    `json:"backgroundUrl"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type CreateParam struct {
	Name string `json:"name" validate:"required"`
}

func (p *CreateParam) Valid() error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return uierr.Invalid("name", "home: name is required")
	}
	if len(p.Name) > 100 {
		return uierr.Invalid("name", "home: name is too long")
	}
	if len(p.Name) < 3 {
		return uierr.Invalid("name", "home: name is too short")
	}
	return nil
}

func Create(ctx context.Context, p *CreateParam) (*Home, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}
	user, err := account.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	var home Home
	err = pgctx.QueryRow(ctx, `insert into home (id, name, user_id, background_url, created_at, updated_at) values ($1, $2, $3, $4, $5, $5) returning id`,
		snowid.Gen(),
		p.Name,
		user.ID,
		"",
		now,
	).Scan(&home.ID)
	if pgctx.UniqueViolation(err, "home_name_user_id_unique") {
		return nil, uierr.Invalid("name", "home: name is already exist")
	}
	if err != nil {
		return nil, err
	}
	return &home, nil
}
