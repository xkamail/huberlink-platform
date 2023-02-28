package account

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrNotFound = uierr.NotFound("user not found")
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

// Find user by id
func Find(ctx context.Context, userID int64) (*User, error) {
	rows, err := pgctx.Query(ctx, ``)
	if err != nil {
		return nil, err
	}
	pgx.CollectOneRow(rows, pgx.RowToStructByPos[User])

}
