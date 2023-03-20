package irremote_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/device"
	"github.com/xkamail/huberlink-platform/iot/irremote"
	"github.com/xkamail/huberlink-platform/pkg/tm"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

//

func TestIRRemote(t *testing.T) {
	m, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	assert.NoError(t, m.CreateTable())
	defer m.Cleanup()

	ctx := tm.CreateUserCtx(t, m.Ctx())

	deviceID, err := device.Create(ctx, &device.CreateParam{})
	assert.NoError(t, err)
	assert.NotNil(t, deviceID)
	remote, err := irremote.Find(ctx, *deviceID)
	assert.NoError(t, err)
	assert.NotNil(t, remote)
	// ordering test
	t.Run("create virtual", func(t *testing.T) {
		var p irremote.CreateVirtualKeyParam
		p.Name = "Air"
		irremote.CreateVirtual(ctx, &p)
	})
	t.Run("list virtual", func(t *testing.T) {
		vs, err := irremote.ListVirtual(ctx, remote.ID)
		assert.NoError(t, err)
		assert.NotNil(t, vs)
		assert.Len(t, vs, 1)
	})
}

func TestVirtualCategory(t *testing.T) {
	t.Parallel()
	t.Run("invalid ", func(t *testing.T) {
		err := json.Unmarshal([]byte(`{"name": "Air","kind":111}`), &irremote.CreateVirtualKeyParam{})
		assert.Error(t, err)
		var uiErr *uierr.Error
		assert.ErrorAs(t, err, &uiErr)
		assert.Equal(t, uierr.CodeInvalidRequest, uiErr.Code)
	})
	t.Run("correct tv ", func(t *testing.T) {
		var p irremote.CreateVirtualKeyParam
		err := json.Unmarshal([]byte(`{"name": "TV","kind":1}`), &p)
		assert.NoError(t, err)
		assert.Equal(t, irremote.VirtualCategoryTV, p.Kind)
	})
	t.Run("correct air ", func(t *testing.T) {
		var p irremote.CreateVirtualKeyParam
		err := json.Unmarshal([]byte(`{"name": "Air","kind":2}`), &p)
		assert.NoError(t, err)
		assert.Equal(t, irremote.VirtualCategoryAirConditioner, p.Kind)
	})
}
