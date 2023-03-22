package irremote

import (
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrVirtualKeyNotFound          = uierr.NotFound("virtual key not found")
	ErrNotFound                    = uierr.NotFound("remote not found")
	ErrCommandNotfound             = uierr.NotFound("command not found")
	ErrRemoteNotFound              = uierr.NotFound("ir remote: remote not found")
	ErrNotLearning                 = uierr.Alert("ir remote: device is not learning mode")
	ErrAnotherVirtualKeyIsLearning = uierr.AlreadyExist("another virtual device is learning please disable it first")
)
