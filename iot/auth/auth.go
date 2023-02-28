package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/pkg/api"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/discord"
	"github.com/xkamail/huberlink-platform/pkg/pgctx"
	"github.com/xkamail/huberlink-platform/pkg/rand"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

// 7 Days
const _refreshTokenLifetime = 24 * time.Hour * 7

type Service struct {
	db        *pgxpool.Pool
	jwtSecret string
	discord   discord.Client
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

func SignInWithDiscord(ctx context.Context, discord discord.Client, code string) (*TokenResponse, error) {
	code = strings.TrimSpace(code)
	if code == "" {
		return nil, uierr.BadInput("code", "code is required")
	}
	accessToken, err := discord.GetAccessToken(ctx, code)
	if err != nil {
		return nil, err
	}
	profile, err := discord.GetProfile(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	var (
		userID int64
	)
	err = pgctx.QueryRow(ctx, `select id from users where discord_id = $1`, profile.ID).Scan(&userID)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	tx, err := pgctx.Begin(ctx)
	defer tx.Rollback(ctx)
	if err != nil {
		return nil, err
	}

	// create a new account
	if userID == 0 {
		now := time.Now()
		err = tx.QueryRow(ctx, `insert into users (id, username, email, password, discord_id, created_at, updated_at) values ($1,$2,$3,$4,$5,$6) returning id`,
			snowid.Gen(),
			profile.Username,
			"",
			profile.ID,
			now,
			now,
		).Scan(&userID)
		if pgctx.UniqueViolation(err, "users_email_unique") {
			return nil, ErrEmailAlreadyExists
		}
		if err != nil {
			return nil, err
		}
	}
	cfg := config.Load()
	jwtToken, err := jwtGenerate(userID, time.Hour*3, cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := createRefreshToken(ctx, tx, userID)
	if err != nil {

		return nil, err
	}
	_, err = tx.Exec(ctx, `update users set updated_at = now() where id = $2`, time.Now(), userID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return &TokenResponse{
		jwtToken,
		refreshToken,
	}, nil
}

func InvokeRefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	var (
		userID    int64
		expiredAt time.Time
	)
	err := pgctx.QueryRow(ctx, `select user_id, expired_at from users_refresh_tokens where token = $1`,
		refreshToken,
	).Scan(
		&userID,
		&expiredAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrRefreshTokenNotFound
	}
	if err != nil {
		return nil, err
	}
	// do generate jwt

	return nil, nil
}

func createRefreshToken(ctx context.Context, tx pgx.Tx, userID int64) (string, error) {
	refreshToken, err := rand.String(300)
	if err != nil {
		return "", err
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
	return refreshToken, err
}

func SignInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("Authorization")
		if len(raw) == 0 {
			api.WriteError(w, ErrNoJWTToken)
			return
		}
		if len(raw) < len("Bearer ") {
			api.WriteError(w, ErrNoJWTToken)
			return
		}
		var jwtToken string
		if len(raw) > len("Bearer ") {
			jwtToken = raw[len("Bearer "):]
		}
		claims, err := jwtVerify(jwtToken)
		if err != nil {
			api.WriteError(w, err)
			return
		}
		acc, err := account.Find(r.Context(), claims.UserID)
		if err != nil {
			api.WriteError(w, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(account.NewContext(r.Context(), acc)))
	})
}
