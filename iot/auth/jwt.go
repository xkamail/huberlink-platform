package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	jwt.StandardClaims
	UserID int64
}

func jwtGenerate(userID int64, expiredAt time.Duration, secret string) (string, error) {
	s, err := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiredAt).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserID: userID,
	}).SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func jwtVerify(token string, secret string) (*JWTClaims, error) {
	p, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr); jwtErr.Errors&jwt.ValidationErrorExpired == jwt.ValidationErrorExpired {
			return nil, ErrTokenExpired
		}
		return nil, ErrJwtParseError
	}
	claims, ok := p.Claims.(*JWTClaims)
	if !ok {
		return nil, ErrJwtParseError
	}
	return claims, nil
}
