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

// Create  a new home and add to member
// need auth
func Create(ctx context.Context, p *CreateParam) (*snowid.ID, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}
	user, err := account.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	tx, err := pgctx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var home Home
	err = pgctx.QueryRow(ctx, `insert into home (id, name, user_id, background_url, created_at, updated_at) values ($1, $2, $3, $4, $5, $5) returning id`,
		snowid.Gen(),
		p.Name,
		user.ID,
		"",
		now,
	).Scan(&home.ID)
	if pgctx.UniqueViolation(err, "home_name_user_id_unique") {
		return nil, uierr.AlreadyExist("home: name is already exist")
	}
	if err != nil {
		return nil, err
	}

	// add member
	_, err = tx.Exec(ctx, `insert into home_members (id, home_id, user_id, permission, created_at, updated_at) values ($1, $2, $3, $4, $5,$5)`,
		snowid.Gen(),
		home.ID,
		user.ID,
		MemberPermissionOwner, // owner
		now,
	)
	if err != nil {
		return nil, err
	}

	// commit the transaction
	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &home.ID, nil
}
