package home_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/iot/home"
	"github.com/xkamail/huberlink-platform/pkg/rand"
	"github.com/xkamail/huberlink-platform/pkg/tm"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

func TestCreate(t *testing.T) {
	testmicro, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testmicro.Cleanup()
	assert.NoError(t, testmicro.CreateTable())
	ctx := testmicro.Ctx()

	// setup user authentication
	userID, err := account.Create(ctx, &account.User{
		Username:  "robert",
		Email:     "some@email.com",
		Password:  "",
		DiscordID: 0,
	})
	assert.NoError(t, err)
	user, err := account.Find(ctx, int64(userID))
	assert.NoError(t, err)
	ctxAuth := account.NewContext(ctx, user)
	assert.NotNilf(t, ctxAuth, "ctxAuth is nil")

	t.Parallel()
	t.Run("success", func(t *testing.T) {

		result, err := home.Create(ctxAuth, &home.CreateParam{
			Name: "test",
		})
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotEqualf(t, 0, *result, "ID is 0")

		t.Run("duplicate", func(t *testing.T) {
			_, err := home.Create(ctxAuth, &home.CreateParam{
				Name: "test",
			})
			assert.Error(t, err)
			var uiErr uierr.Error
			assert.ErrorAs(t, err, &uiErr)
			assert.Equal(t, uierr.CodeAlreadyExists, uiErr.Code)
		})
	})

	t.Run("length too short", func(t *testing.T) {
		_, err := home.Create(ctxAuth, &home.CreateParam{
			Name: "t",
		})
		assert.Error(t, err)
	})
	t.Run("length too long", func(t *testing.T) {
		name, _ := rand.String(101)
		_, err := home.Create(ctxAuth, &home.CreateParam{
			Name: name,
		})
		assert.Error(t, err)
	})
}
