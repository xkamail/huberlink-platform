package device

import (
	"encoding/json"

	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var ErrInvalidKind = uierr.Invalid("kind", "invalid kind")

type Kind uint

func (k *Kind) UnmarshalJSON(b []byte) error {
	// unmarshal uint
	var i uint
	if err := json.Unmarshal(b, &i); err != nil {
		return err
	}
	if i > uint(KindLamp) {
		return ErrInvalidKind
	}
	*k = Kind(i)
	return nil
}

func (k *Kind) MarshalJSON() ([]byte, error) {
	// de-pointer to get a value
	return json.Marshal(uint(*k))
}

const (
	KindUnknown Kind = iota
	KindIRRemote
	KindSensor
	KindSwitch
	KindLamp
)

var _ json.Marshaler = (*Kind)(nil)
var _ json.Unmarshaler = (*Kind)(nil)
