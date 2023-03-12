package auth

import (
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

var (
	ErrTokenExpired                 = uierr.New(uierr.CodeTokenExpired, "token expired")
	ErrJwtParseError                = uierr.Alert("JWT parse error")
	ErrJwtInvalidToken              = uierr.Alert("JWT invalid token")
	ErrRefreshTokenNotFound         = uierr.NotFound("refresh token not found")
	ErrUsernameAndPasswordIncorrect = uierr.Alert("Username or password is incorrect")
	ErrNoJWTToken                   = uierr.UnAuthorization("no authorization token")
	ErrInvalidJWTSchema             = uierr.BadInput("jwt", "invalid jwt schema")
	ErrEmailAlreadyExists           = uierr.AlreadyExist("email address already exists")
	ErrRefreshTokenExpired          = uierr.Alert("refresh token expired")
	ErrPasswordAlreadySet           = uierr.Alert("password already set")
)
