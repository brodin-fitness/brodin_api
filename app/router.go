package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Initializes all routes and middleware.
func initRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	return r
}
