package irremote

import (
	"context"
	"time"

	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

// VirtualCategory is an enum that represents a virtual key category
type VirtualCategory uint

const (
	VirtualCategoryOther VirtualCategory = iota
	VirtualCategoryTV
	VirtualCategoryAirConditioner
	VirtualCategoryLight
	VirtualCategoryFan
	VirtualCategorySpeaker
	VirtualCategoryProjector
	VirtualCategoryDVD
	VirtualCategoryWaterHeater
)

// IRRemote is a struct that represents a remote control
// that linked to a device.Device and home.Home
type IRRemote struct {
	ID        snowid.ID `json:"id"`
	DeviceID  snowid.ID `json:"deviceId"`
	HomeID    snowid.ID `json:"homeId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// VirtualKey is a struct that represents a remote control
// virtual remote control
type VirtualKey struct {
	ID       snowid.ID       `json:"id"`
	RemoteID snowid.ID       `json:"remoteId"`
	Name     string          `json:"name"`
	Icon     string          `json:"icon"`
	Kind     VirtualCategory `json:"category"`
	// IsLearning is a flag that indicates that the virtual remote is learning
	// when rawData codes has come
	// Command will be created and IsLearning will be false
	//
	IsLearning bool      `json:"isLearning"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Command struct {
	ID        snowid.ID `json:"id"`
	RemoteID  snowid.ID `json:"remoteId"`
	VirtualID snowid.ID `json:"virtualId"`
	Name      string    `json:"name"`
	Code      []uint    `json:"code"`
	Remark    *string   `json:"remark"`
	Platform  string    `json:"platform"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func Find(ctx context.Context, deviceID snowid.ID) (*IRRemote, error) {
	// TODO: implement
	panic("not implemented")
}

func ListVirtual(ctx context.Context, deviceID snowid.ID) ([]*VirtualKey, error) {
	// TODO: implement
	panic("not implemented")
}

type CreateVirtualKeyParam struct {
}

func CreateVirtual(ctx context.Context, p *CreateVirtualKeyParam) (*VirtualKey, error) {
	// TODO: implement
	panic("not implemented")
}
