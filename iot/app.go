package iot

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/iot/auth"
	"github.com/xkamail/huberlink-platform/iot/home"
	"github.com/xkamail/huberlink-platform/pkg/api"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/discord"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
	"github.com/xkamail/snowflake"
)

type Validator interface {
	Valid() error
}

func Handlers() http.Handler {
	cfg := config.Load()

	discordClient := discord.NewClient(cfg.DiscordClientID, cfg.DiscordClientSecret, cfg.DiscordRedirectUri)

	router := chi.NewRouter()
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		api.WriteError(w, uierr.NotFound("api entry not found"))
	})
	authRouter := router.With(auth.SignInMiddleware) // auth middleware

	// auth
	{
		router.Post("/auth/sign-in-username", h(func(ctx context.Context, r *http.Request) (any, error) {
			var p auth.SignInParam
			if err := mustBind(r, &p); err != nil {
				return nil, err
			}
			return auth.SignIn(ctx, &p)
		}))
		router.Post("/auth/sign-in", h(func(ctx context.Context, r *http.Request) (any, error) {
			var p auth.SignInWithDiscordParam
			if err := mustBind(r, &p); err != nil {
				return nil, err
			}
			return auth.SignInWithDiscord(ctx, discordClient, &p)
		}))
		router.Post("/auth/refresh-token", h(func(ctx context.Context, r *http.Request) (any, error) {

			code := r.URL.Query().Get("refreshToken")
			return auth.InvokeRefreshToken(ctx, code)
		}))
		authRouter.Get("/auth/me", h(func(ctx context.Context, r *http.Request) (any, error) {

			return account.FromContext(ctx)
		}))
	}
	// user
	{
		authRouter.Get("/user/me", nil)
	}
	// home
	{
		// list my home
		authRouter.Get("/home", h(func(ctx context.Context, r *http.Request) (any, error) {
			acc, err := account.FromContext(ctx)
			if err != nil {
				return nil, err
			}
			return home.List(ctx, acc.ID)
		}))
		// create home
		authRouter.Post("/home", h(func(ctx context.Context, r *http.Request) (any, error) {
			var p home.CreateParam
			if err := mustBind(r, &p); err != nil {
				return nil, err
			}
			return home.Create(ctx, &p)
		}))

		r := authRouter.With(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				acc, err := account.FromContext(r.Context())
				if err != nil {
					api.WriteError(w, err)
					return
				}
				homeID, err := URLParamID(r, "home_id")
				if err != nil {
					api.WriteError(w, err)
					return
				}
				h, err := home.GetFromIDAndUserID(r.Context(), homeID, acc.ID)
				if err != nil {
					api.WriteError(w, err)
					return
				}
				next.ServeHTTP(w, r.WithContext(newHomeContext(r.Context(), h)))
			})
		})
		// get detail home
		r.Get("/home/{home_id}", h(func(ctx context.Context, r *http.Request) (any, error) {
			h, err := homeFromCtx(ctx)
			if err != nil {
				return nil, err
			}
			var result struct {
				Home    *home.Home     `json:"home"`
				Members []*home.Member `json:"members"`
			}
			result.Home = h

			member, err := home.ListMember(ctx, h.ID)
			if err != nil {
				return nil, err
			}
			result.Members = member

			return result, nil
		}))
		// edit home info
		r.Put("/home/{id}", h(func(ctx context.Context, r *http.Request) (any, error) {
			h, err := homeFromCtx(ctx)
			if err != nil {
				return nil, err
			}
			_ = h
			panic("not implemented")
			// TODO: implement
		}))
		// join home
	}
	// device
	{
		authRouter.Get("/home/{home_id}/devices/all", nil)
		authRouter.Post("/home/{home_id}/devices", nil)
		authRouter.Get("/home/{home_id}/devices/{id}", nil)
		authRouter.Delete("/home/{home_id}/devices/{id}", nil)
		authRouter.Patch("/home/{home_id}/devices/{id}", nil)
	}
	return router
}

func h(fn func(ctx context.Context, r *http.Request) (any, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := fn(r.Context(), r)
		if err != nil {
			api.WriteError(w, err)
			return
		}
		api.Write(w, res)
	}
}

// mustBind a json to a struct
// return error when invalid
func mustBind(r *http.Request, v any) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	if vv, ok := v.(interface{ Valid() error }); ok {
		if err := vv.Valid(); err != nil {
			return err
		}
	}
	return nil
}

func URLParamID(r *http.Request, key string) (snowid.ID, error) {
	id := chi.URLParam(r, key)
	i, err := snowflake.ParseString(id)
	if err != nil {
		return snowid.Zero, uierr.Invalid("id", "invalid id parameter")
	}
	return snowid.ID(i), nil
}

// context key use on home detail api route
type homeDetailCtx struct {
}

func newHomeContext(ctx context.Context, h *home.Home) context.Context {
	return context.WithValue(ctx, homeDetailCtx{}, h)
}

func homeFromCtx(ctx context.Context) (*home.Home, error) {
	h, ok := ctx.Value(homeDetailCtx{}).(*home.Home)
	if !ok {
		return nil, uierr.UnAuthorization("home: home not found")
	}
	return h, nil
}
