package irremote

import (
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var ErrVirtualKeyNotfound = uierr.NotFound("virtual key not found")
var ErrNotFound = uierr.NotFound("remote not found")

var ErrCommandNotfound = uierr.NotFound("command not found")
