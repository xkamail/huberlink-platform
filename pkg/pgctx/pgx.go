package pgctx

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ctxKey struct {
}

func NewContext(ctx context.Context, pool *pgxpool.Pool) context.Context {
	return context.WithValue(ctx, ctxKey{}, pool)
}

func Middleware(db *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(NewContext(r.Context(), db)))
		})
	}
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

func Collect[T any](ctx context.Context, sql string, args ...any) ([]*T, error) {
	db := q(ctx)
	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	result, err := pgx.CollectRows(rows, pgx.RowToStructByPos[*T])
	if err != nil {
		// create empty array
		return make([]*T, 0), err
	}
	return result, nil
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
	return q(ctx).BeginTx(ctx, txOptions)
}
