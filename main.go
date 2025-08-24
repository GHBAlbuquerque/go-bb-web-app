package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf("%s", fmt.Sprintf("Starting server on port %s...\n", portNumber))
	http.ListenAndServe(portNumber, nil)
}
