package main

import (
	"net/http"

	"github.com/GHBAlbuquerque/go-bb-web-app/pkg/config"
	"github.com/GHBAlbuquerque/go-bb-web-app/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux = a multiplexer
	// It comes from networking/telecom: a multiplexer takes many input signals and funnels them into one output channel (or vice versa).
	mux := chi.NewRouter()

	// Middleware = functions that wrap your handlers to add common behavior before/after each request.
	mux.Use(middleware.Logger, middleware.Recoverer, NoSurf, SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/rooms/general-quarters", handlers.Repo.GeneralsQuarters)
	mux.Get("/rooms/marjors-suite", handlers.Repo.MarjorsSuite)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
