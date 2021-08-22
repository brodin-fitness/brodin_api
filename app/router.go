package app

import (
	"github.com/brodin_fitness/brodin_api/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Initializes all routes and middleware.
func initRouter() *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", controllers.Ping)

	return r
}
