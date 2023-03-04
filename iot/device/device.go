package device

import (
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

type Kind uint

const (
	KindUnknown Kind = iota
	KindIRRemote
)

type Device struct {
	ID                snowid.ID `json:"id"`
	Name              string    `json:"name"`
	Icon              string    `json:"icon"`
	Model             string    `json:"model"`
	Kind              Kind      `json:"kind"`
	HomeID            snowid.ID `json:"homeId"`
	UserID            snowid.ID `json:"userId"`
	Token             string    `json:"token"`
	IpAddress         string    `json:"ipAddress"`
	Location          *string   `json:"location"`
	LatestHeartbeatAt string    `json:"latestHeartbeatAt"`
	CreatedAt         string    `json:"createdAt"`
	UpdatedAt         string    `json:"updatedAt"`
}
