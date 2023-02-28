package account

import (
	"context"

	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var ErrNoUserContext = uierr.UnAuthorization("no user in request context")

type ctxKey struct {
}

func NewContext(ctx context.Context, acc *User) context.Context {
	return context.WithValue(ctx, ctxKey{}, acc)
}

func FromContext(ctx context.Context) (*User, error) {
	acc, b := ctx.Value(ctxKey{}).(*User)
	if !b {
		return nil, ErrNoUserContext
	}
	return acc, nil
}
