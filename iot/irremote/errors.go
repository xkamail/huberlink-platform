package irremote

import (
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrVirtualKeyNotfound = uierr.NotFound("virtual key not found")
	ErrNotFound           = uierr.NotFound("remote not found")
	ErrCommandNotfound    = uierr.NotFound("command not found")
	ErrRemoteNotFound     = uierr.NotFound("ir remote: remote not found")
)
