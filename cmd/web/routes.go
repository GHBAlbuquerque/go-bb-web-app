package main

import (
	"go-bb-web-app/pkg/config"
	"go-bb-web-app/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux = a multiplexer
	// It comes from networking/telecom: a multiplexer takes many input signals and funnels them into one output channel (or vice versa).
	mux := chi.NewRouter()

	// Middleware = functions that wrap your handlers to add common behavior before/after each request.
	mux.Use(middleware.Logger, middleware.Recoverer, WriteToConsole, NoSurf, SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
