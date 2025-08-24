package main

import (
	"fmt"
	"go-bb-web-app/pkg/handlers"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("%s", fmt.Sprintf("Starting server on port %s...\n", portNumber))
	http.ListenAndServe(portNumber, nil)
}
