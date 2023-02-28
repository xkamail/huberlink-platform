package pgctx

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ctxKey struct {
}

func NewContext(ctx context.Context, pool *pgxpool.Pool) context.Context {
	return context.WithValue(ctx, ctxKey{}, pool)
}

func q(ctx context.Context) *pgxpool.Pool {
	pool, b := ctx.Value(ctxKey{}).(*pgxpool.Pool)
	if !b {
		panic("no pgxpool context")
	}
	return pool
}
func Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return q(ctx).Query(ctx, sql, args...)
}

func QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return q(ctx).QueryRow(ctx, sql, args...)
}
func Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	return q(ctx).Exec(ctx, sql, args...)
}

func Begin(ctx context.Context) (pgx.Tx, error) {
	return q(ctx).Begin(ctx)
}

func BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return q(ctx).BeginTx(ctx)
}
