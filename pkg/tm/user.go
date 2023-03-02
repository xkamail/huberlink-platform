package tm

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/rand"
)

func CreateUserCtx(t *testing.T, ctx context.Context) context.Context {
	name, _ := rand.String(10)
	// setup user authentication
	userID, err := account.Create(ctx, &account.User{
		Username:  name,
		Email:     fmt.Sprintf("%s@%s.com", name, name),
		Password:  "",
		DiscordID: 0,
	})
	assert.NoError(t, err)
	user, err := account.Find(ctx, int64(userID))
	assert.NoError(t, err)
	ctxAuth := account.NewContext(ctx, user)
	assert.NotNilf(t, ctxAuth, "ctxAuth is nil")
	return ctxAuth
}
