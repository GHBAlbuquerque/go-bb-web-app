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
	mux.Use(middleware.Logger, middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

/*
Could have been done with just the standard lib:

func routes() http.Handler {
	mux := http.NewServeMux()

	// simple routes
	mux.HandleFunc("GET /", handlers.Repo.Home)
	mux.HandleFunc("GET /about", handlers.Repo.About)

	// example with path param and method
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // stdlib path param
		w.Write([]byte("user id = " + id))
	})

	return mux
}

func main() {
	http.ListenAndServe(":8080", routes())
}
*/
