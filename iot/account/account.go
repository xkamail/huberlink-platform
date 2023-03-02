package account

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrNotFound          = uierr.NotFound("user not found")
	ErrEmailDuplicate    = uierr.AlreadyExist("email address already exists")
	ErrUsernameDuplicate = uierr.AlreadyExist("username already exists")
)

const minUsernameLength = 5

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

func FindByUsername(ctx context.Context, username string) (*User, error) {
	rows, err := pgctx.Query(ctx, `select id, username, email, password, discord_id, created_at, updated_at from users where username = $1`, username)
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

func FindByEmail(ctx context.Context, email string) (*User, error) {
	rows, err := pgctx.Query(ctx, `select id, username, email, password, discord_id, created_at, updated_at from users where email = $1`, email)
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

func Create(ctx context.Context, user *User) (snowid.ID, error) {
	if len(user.Username) < minUsernameLength {
		return 0, uierr.BadInput("username", fmt.Sprintf("username must be at least %d characters", minUsernameLength))
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return 0, uierr.BadInput("email", "invalid email address")
	}
	now := time.Now()
	err := pgctx.QueryRow(ctx, `
			insert into users (id, username, email, password, discord_id, created_at, updated_at)
			values ($1, $2, $3, $4, $5, $6, $7) returning id`,
		snowid.Gen(),
		strings.TrimSpace(user.Username),
		user.Email,
		user.Password,
		user.DiscordID,
		now,
		now,
	).Scan(&user.ID)
	if pgctx.UniqueViolation(err, "users_username_unique") {
		return 0, ErrUsernameDuplicate
	}
	if pgctx.UniqueViolation(err, "users_email_unique") {
		return 0, ErrEmailDuplicate
	}
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
