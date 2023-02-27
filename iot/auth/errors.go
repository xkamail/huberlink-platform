package auth

import (
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrTokenExpired                 = uierr.New(uierr.CodeTokenExpired, "token expired")
	ErrJwtParseError                = uierr.Alert("JWT parse error")
	ErrJwtInvalidToken              = uierr.Alert("JWT invalid token")
	ErrRefreshTokenNotFound         = uierr.NotFound("Refresh token not found")
	ErrUsernameAndPasswordIncorrect = uierr.Alert("Username or password is incorrect")
)
