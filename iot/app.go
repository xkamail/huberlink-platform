package iot

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handlers() http.Handler {
	router := chi.NewRouter()
	// auth
	{
		router.Post("/auth/sign-in", nil)
		router.Post("/auth/refresh-token", nil)
		router.Get("/auth/me", nil)
	}
	// user
	{
		router.Get("/user/me", nil)
	}
	// home
	return router
}
