package irremote_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xkamail/huberlink-platform/iot/device"
	"github.com/xkamail/huberlink-platform/pkg/tm"
)

//

func TestIRRemote(t *testing.T) {
	m, err := tm.New(t, "../../db/migrations")
	assert.NoError(t, err)
	assert.NoError(t, m.CreateTable())
	defer m.Cleanup()

	ctx := tm.CreateUserCtx(t, m.Ctx())
	//
	t.Run("create ir remote device", func(t *testing.T) {
		create, err := device.Create(ctx, &device.CreateParam{})
		assert.NoError(t, err)
		assert.NotNil(t, create)
	})
}
