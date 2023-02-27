package auth

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/xkamail/huberlink-platform/pkg/rand"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
)

// 7 Days
const _refreshTokenLifetime = 24 * time.Hour * 7

type Service struct {
	db        *pgxpool.Pool
	jwtSecret string
}

func NewService(db *pgxpool.Pool, jwtSecret string) *Service {
	return &Service{db, jwtSecret}
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func (s Service) SignInWithDiscord(ctx context.Context) (*TokenResponse, error) {

	return nil, nil
}

func (s Service) InvokeRefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	var (
		userID    int64
		expiredAt time.Time
	)
	err := s.db.QueryRow(ctx, `select user_id, expired_at from users_refresh_tokens where token = $1`, refreshToken).Scan(&userID, &expiredAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrRefreshTokenNotFound
	}
	if err != nil {
		return nil, err
	}
	// do generate jwt

	return nil, nil
}

func (s Service) createRefreshToken(ctx context.Context, tx pgxpool.Tx, userID int64) error {
	refreshToken, err := rand.String(300)
	if err != nil {
		return err
	}
	now := time.Now()
	_, err = tx.Exec(ctx, `
		insert into users_refresh_tokens 
		    (id,user_id,token,expired_at,issued_at,created_at) 
		values ($1,$2,$3,$4,$5,$6)`,
		snowid.Gen(),
		userID,
		refreshToken,
		now.Add(_refreshTokenLifetime),
		now,
		now,
	)
	return err
}
