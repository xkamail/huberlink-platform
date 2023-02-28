package iot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/iot/auth"
	"github.com/xkamail/huberlink-platform/pkg/api"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/discord"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
)

func Handlers() http.Handler {
	cfg := config.Load()

	discordClient := discord.NewClient(cfg.DiscordClientID, cfg.DiscordClientSecret, cfg.DiscordRedirectURI)

	router := chi.NewRouter()
	// auth
	{
		router.Post("/auth/sign-in", h(func(ctx context.Context, r *http.Request) (*auth.TokenResponse, error) {

			code := r.URL.Query().Get("code")
			return auth.SignInWithDiscord(r.Context(), discordClient, code)
		}))
		router.Post("/auth/refresh-token", h(func(ctx context.Context, r *http.Request) (*auth.TokenResponse, error) {

			code := r.URL.Query().Get("refreshToken")
			return auth.InvokeRefreshToken(r.Context(), code)
		}))
		router.With(auth.SignInMiddleware).Get("/auth/me", h(func(ctx context.Context, r *http.Request) (any, error) {

			return account.FromContext(ctx)
		}))
	}
	// user
	{
		router.Get("/user/me", nil)
	}
	// home
	return router
}

func h[T any](fn func(ctx context.Context, r *http.Request) (T, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := fn(r.Context(), r)
		if err != nil {
			message := err.Error()
			errs := make([]any, 0)
			var errCode uierr.Code

			if uiErr, ok := err.(*uierr.Error); ok {
				errCode = uiErr.Code()
				errs = append(errs, uiErr)
				message = uiErr.Message()
			}

			_ = json.NewEncoder(w).Encode(&api.Format{
				Success: false,
				Data:    nil,
				Errors:  errs,
				Message: message,
				Code:    errCode,
			})
			return
		}

		_ = json.NewEncoder(w).Encode(&api.Format{
			Success: true,
			Data:    t,
			Errors:  nil,
			Message: "Success",
		})

	}
}
