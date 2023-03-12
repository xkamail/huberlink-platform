package iot

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/xkamail/snowflake"

	"github.com/xkamail/huberlink-platform/iot/account"
	"github.com/xkamail/huberlink-platform/iot/auth"
	"github.com/xkamail/huberlink-platform/iot/device"
	"github.com/xkamail/huberlink-platform/iot/home"
	"github.com/xkamail/huberlink-platform/pkg/api"
	"github.com/xkamail/huberlink-platform/pkg/config"
	"github.com/xkamail/huberlink-platform/pkg/discord"
	"github.com/xkamail/huberlink-platform/pkg/snowid"
	"github.com/xkamail/huberlink-platform/pkg/uierr"
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
		authRouter.Put("/auth/set-password", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				var p auth.SetPasswordParam
				if err := mustBind(r, &p); err != nil {
					return nil, err
				}
				if err := auth.SetPassword(ctx, &p); err != nil {
					return nil, err
				}
				return true, nil
			},
		))
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

		// middleware get home from url param
		// and validate user is member of home
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
		// device
		r.Get("/home/{home_id}/devices/all", h(func(ctx context.Context, r *http.Request) (any, error) {
			h, err := homeFromCtx(ctx)
			if err != nil {
				return nil, err
			}
			return device.List(ctx, h.ID)
		}))
		r.Post("/home/{home_id}/devices", h(func(ctx context.Context, r *http.Request) (any, error) {
			h, err := homeFromCtx(ctx)
			if err != nil {
				return nil, err
			}
			var p device.CreateParam
			if err := mustBind(r, &p); err != nil {
				return nil, err
			}
			p.HomeID = h.ID
			return device.Create(ctx, &p)
		}))
		r.Get("/home/{home_id}/devices/{device_id}", h(func(ctx context.Context, r *http.Request) (any, error) {
			deviceID, err := URLParamID(r, "device_id")
			if err != nil {
				return nil, err
			}
			return device.Find(ctx, deviceID)
		}))
		r.Delete("/home/{home_id}/devices/{id}", nil)
		r.Patch("/home/{home_id}/devices/{id}", nil)
		// ir-remote service
		r.Get("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// return list of virtual remote
				return nil, errors.New("not implemented")
			},
		))
		// create virtual remote
		r.Post("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// define a virtual remote
				// kind
				return nil, errors.New("not implemented")
			},
		))
		r.Put("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// update virtual remote information and kind ?
				//
				return nil, errors.New("not implemented")
			},
		))
		r.Delete("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// delete cascade all button
				// which in this virtual remote
				return nil, errors.New("not implemented")
			},
		))
		r.Get("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// return a list of button
				return nil, errors.New("not implemented")
			},
		))
		r.Post("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}/start-learning", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// create a 30 seconds learning session
				// when universal remote got an ir signal
				// then will create a button
				return nil, errors.New("not implemented")
			},
		))
		r.Post("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}/execute", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// execute a button
				return nil, errors.New("not implemented")
			},
		))
		r.Put("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}/button/{button_id}", h(
			func(ctx context.Context, r *http.Request) (any, error) {
				// if name is empty then it has to be a new remote button
				// which come from learning session
				// update a button name of a virtual remote
				return nil, errors.New("not implemented")
			},
		))
		r.Delete("/home/{home_id}/devices/{devices_id}/ir-remote/{remote_id}/virtual/{virtual_id}/button/{button_id}", h(func(ctx context.Context, r *http.Request) (any, error) {
			// delete a button and codes is gone
			return nil, errors.New("not implemented")
		}))
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
