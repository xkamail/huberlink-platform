package irremote

import (
	"context"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

func StartLearning(ctx context.Context, deviceID, virtualID snowid.ID) error {
	// check if any virtual key is learning then return error
	// to prevent multiple learning
	var learningCount uint64
	err := pgctx.QueryRow(ctx, `
		select count(*) 
			from device_ir_remote_virtual_keys
			inner join device_ir_remotes dir on device_ir_remote_virtual_keys.remote_id = dir.id
			where is_learning = true and dir.device_id = $1`,
		deviceID,
	).Scan(&learningCount)
	if err != nil {
		return err
	}
	if learningCount > 0 {
		return ErrAnotherVirtualKeyIsLearning
	}
	affect, err := pgctx.Exec(ctx, `update device_ir_remote_virtual_keys set is_learning = true where id = $1`, virtualID)
	if err != nil {
		return err
	}
	if affect.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func StopLearning(ctx context.Context, deviceID, virtualID snowid.ID) error {
	// stop learning
	// to disable learning modal on frontend
	_, err := pgctx.Exec(ctx, `update device_ir_remote_virtual_keys set is_learning = false where id = $1`)
	return err
}
