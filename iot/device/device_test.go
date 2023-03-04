package device_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/device"
	"github.com/xkamail/huberlink-platform/iot/home"
	"github.com/xkamail/huberlink-platform/pkg/tm"
)

func TestCreate(t *testing.T) {
	// testm
	testm, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	defer testm.Close()
	assert.NoError(t, testm.CreateTable())
	ctx := testm.Ctx()
	authCtx := tm.CreateUserCtx(t, ctx)
	homeID, err := home.Create(authCtx, &home.CreateParam{
		Name: "condo",
	})
	assert.NoError(t, err)
	assert.NotNil(t, homeID)
	t.Parallel()
	t.Run("no auth should error", func(t *testing.T) {
		_, err := device.Create(ctx, &device.CreateParam{})
		assert.Error(t, err)
		t.Log(err)
	})
	t.Run("create device", func(t *testing.T) {
		dev, err := device.Create(authCtx, &device.CreateParam{
			Name:   "ir-remote-1",
			Model:  "ZA117",
			Kind:   device.KindIRRemote,
			HomeID: *homeID,
		})
		assert.NoError(t, err)
		assert.NotNil(t, dev)
		t.Log(*dev)
	})

}
