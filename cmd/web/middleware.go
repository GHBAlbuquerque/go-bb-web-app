package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WriteToConsole is a custom middleware created just for testing
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page!")
		next.ServeHTTP(w, r)
	})
}

// NoSurf creates base cookie for the site using the nosurf lib
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}
