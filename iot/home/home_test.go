package home_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

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
