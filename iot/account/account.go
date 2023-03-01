package account

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrNotFound = uierr.NotFound("user not found")
)

type User struct {
	ID        snowid.ID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	DiscordID snowid.ID `json:"discordId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Find user by id
func Find(ctx context.Context, userID int64) (*User, error) {
	rows, err := pgctx.Query(ctx, `select id, username, email, password, discord_id, created_at, updated_at from users where id = $1`, userID)
	if err != nil {
		return nil, err
	}
	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[User])
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
