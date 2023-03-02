package home_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/iot/home"
	"github.com/xkamail/huberlink-platform/pkg/rand"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/tm"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

func TestList(t *testing.T) {
	testmicro, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testmicro.Cleanup()
	assert.NoError(t, testmicro.CreateTable())
	ctx := testmicro.Ctx()
	ctxAuth := tm.CreateUserCtx(t, ctx)

	user, err := account.FromContext(ctxAuth)
	assert.NoError(t, err)

	for i := 0; i < home.MaxHomePerUser; i++ {
		_, err := home.Create(ctxAuth, &home.CreateParam{
			Name: fmt.Sprintf("home 00%d", i),
		})
		assert.NoError(t, err)
	}

	t.Parallel()
	t.Run("list 10", func(t *testing.T) {
		homes, err := home.List(ctxAuth, user.ID)
		assert.NoError(t, err)
		assert.Equal(t, home.MaxHomePerUser, len(homes))
	})
	t.Run("list user not found", func(t *testing.T) {
		homes, err := home.List(ctxAuth, snowid.Gen())
		assert.Nil(t, err)
		assert.Len(t, homes, 0)
	})

}
func TestCreate(t *testing.T) {
	testmicro, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testmicro.Cleanup()
	assert.NoError(t, testmicro.CreateTable())
	ctx := testmicro.Ctx()

	ctxAuth := tm.CreateUserCtx(t, ctx)

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
		t.Run("another should not be duplicate", func(t *testing.T) {
			ctxAuth2 := tm.CreateUserCtx(t, ctx)
			result, err := home.Create(ctxAuth2, &home.CreateParam{
				Name: "test",
			})
			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.NotEqualf(t, 0, *result, "ID is 0")
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
