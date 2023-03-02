package home

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrReachLimitHome = uierr.Alert("home: reach limit home per user")
	ErrNotFound       = uierr.NotFound("home: not found")
)

const MaxHomePerUser = 5

type Home struct {
	ID            snowid.ID `json:"id"`
	Name          string    `json:"name"`
	UserID        snowid.ID `json:"userId"`
	BackgroundURL string    `json:"backgroundUrl"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func List(ctx context.Context, userID snowid.ID) ([]*Home, error) {
	homes, err := pgctx.Collect[Home](ctx, `select id, name, user_id, background_url, created_at, updated_at from home where user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	return homes, nil
}

func GetFromIDAndUserID(ctx context.Context, homeID, userID snowid.ID) (*Home, error) {
	rows, err := pgctx.Query(ctx, `
		select h.id, h.name, h.user_id, h.background_url, h.created_at, h.updated_at 
		from home h 
		inner join home_members hm 
		on h.id = hm.home_id 
		where hm.user_id = $1 and hm.home_id = $2`,
		userID,
		homeID,
	)
	if err != nil {
		return nil, err
	}
	h, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[Home])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return h, nil
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
// need auth limit 5 home per users
func Create(ctx context.Context, p *CreateParam) (*snowid.ID, error) {
	if err := p.Valid(); err != nil {
		return nil, err
	}
	user, err := account.FromContext(ctx)
	if err != nil {
		return nil, err
	}
	now := time.Now()

	var count int
	err = pgctx.QueryRow(ctx, `select count(*) from home where user_id = $1`, user.ID).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count >= MaxHomePerUser {
		return nil, ErrReachLimitHome
	}

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
