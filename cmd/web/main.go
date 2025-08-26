package main

import (
	"fmt"
	"go-bb-web-app/pkg/config"
	"go-bb-web-app/pkg/handlers"
	"go-bb-web-app/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create Template Cache:", err)
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	render.SetAppConfig(&app)
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("%s", fmt.Sprintf("Starting server on port %s...\n", portNumber))
	http.ListenAndServe(portNumber, nil)
}
