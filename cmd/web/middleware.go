package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf creates base cookie for the site using the nosurf lib
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad automatically loads and saves session data for the current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
