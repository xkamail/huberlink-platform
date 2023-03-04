package account_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/tm"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

func TestCreate(t *testing.T) {
	testmicro, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testmicro.Cleanup()
	assert.NoError(t, testmicro.CreateTable())
	t.Parallel()
	ctx := testmicro.Ctx()
	t.Run("username too short", func(t *testing.T) {
		id, err := account.Create(ctx, &account.User{
			ID:       0,
			Username: "1234",
			Email:    "",
		})
		assert.Error(t, err)
		var uiErr *uierr.Error
		assert.ErrorAsf(t, err, &uiErr, "must be ErrUsernameTooShort")
		assert.Equalf(t, snowid.Zero, id, "id must be zero")
	})
	t.Run("invalid email address", func(t *testing.T) {
		id, err := account.Create(ctx, &account.User{
			ID:       0,
			Username: "123456",
			Email:    "xxx",
		})
		assert.Error(t, err)
		var uiErr *uierr.Error
		assert.ErrorAsf(t, err, &uiErr, "must be ErrUsernameTooShort")
		assert.Equal(t, uierr.CodeInvalidRequest, uiErr.Code)
		assert.Equalf(t, snowid.Zero, id, "id must be zero")
	})
	t.Run("create success", func(t *testing.T) {
		id, err := account.Create(ctx, &account.User{
			ID:        0,
			Username:  "123456",
			Email:     "some@email.com",
			Password:  "",
			DiscordID: 0,
		})
		assert.NoError(t, err)
		assert.NotEqualf(t, 0, id, "id must not be zero")
	})
	t.Run("create with duplicate email", func(t *testing.T) {
		dupEmail := "dup@email.com"
		id, err := account.Create(ctx, &account.User{
			ID:        0,
			Username:  "123456xxx",
			Email:     dupEmail,
			Password:  "",
			DiscordID: 0,
		})
		assert.NoError(t, err)
		assert.NotEqualf(t, 0, id, "id must not be zero")

		id2, err2 := account.Create(ctx, &account.User{
			ID:        0,
			Username:  "123456xxx",
			Email:     dupEmail,
			Password:  "",
			DiscordID: 0,
		})
		assert.Errorf(t, err2, "must be error")
		assert.ErrorIs(t, account.ErrEmailDuplicate, err2)
		assert.Equalf(t, snowid.Zero, id2, "id must be zero")
	})
	t.Run("create with duplicate username", func(t *testing.T) {
		dupUsername := "dupUsername"
		id, err := account.Create(ctx, &account.User{
			ID:        0,
			Username:  dupUsername,
			Email:     "random@email.com",
			Password:  "",
			DiscordID: 0,
		})
		assert.NoError(t, err)
		assert.NotEqualf(t, 0, id, "id must not be zero")

		id2, err2 := account.Create(ctx, &account.User{
			ID:        0,
			Username:  dupUsername,
			Email:     "random2@email.com",
			Password:  "",
			DiscordID: 0,
		})
		assert.Errorf(t, err2, "must be error")
		assert.ErrorIs(t, account.ErrUsernameDuplicate, err2)
		assert.Equalf(t, snowid.Zero, id2, "id must be zero")
	})
}

func TestFind(t *testing.T) {
	testmicro, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testmicro.Cleanup()
	assert.NoError(t, testmicro.CreateTable())
	t.Parallel()
}

func TestFindByUsername(t *testing.T) {
	testmicro, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testmicro.Cleanup()
	assert.NoError(t, testmicro.CreateTable())
	t.Parallel()
}
