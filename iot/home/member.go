package home

import (
	"time"

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
