package main

import (
	"fmt"
	"go-bb-web-app/pkg/config"
	"go-bb-web-app/pkg/handlers"
	"go-bb-web-app/pkg/render"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = createSession()
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Template Cache:", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	render.SetAppConfig(&app)
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	log.Printf("%s", fmt.Sprintf("Starting server on port %s...\n", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

func createSession() *scs.SessionManager {
	newSession := scs.New()
	newSession.Lifetime = 24 * time.Hour
	newSession.Cookie.Persist = true
	newSession.Cookie.SameSite = http.SameSiteLaxMode
	newSession.Cookie.Secure = app.InProduction

	return newSession
}
