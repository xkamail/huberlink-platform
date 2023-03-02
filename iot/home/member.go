package home

import (
	"context"
	"time"

	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

type Member struct {
	ID         snowid.ID        `json:"id"`
	HomeID     snowid.ID        `json:"homeId"`
	UserID     snowid.ID        `json:"userId"`
	Permission MemberPermission `json:"permission"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
}

type MemberPermission uint64

const (
	MemberPermissionGuest MemberPermission = 2 << iota
	MemberPermissionOwner
	MemberPermissionManageDevice
	MemberPermissionManageMember
	MemberPermissionManageHome
	MemberPermissionManageRoom
)

func ListMember(ctx context.Context, homeID snowid.ID) ([]*Member, error) {
	// query from db and return
	collect, err := pgctx.Collect[Member](ctx, `select id, home_id, user_id, permission, created_at, updated_at from home_members where home_id = $1`,
		homeID,
	)
	if err != nil {
		return nil, err
	}
	return collect, nil
}
