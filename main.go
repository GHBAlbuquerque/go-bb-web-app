package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "home.page")
}

func About(writer http.ResponseWriter, request *http.Request) {
	renderTemplate(writer, "about.page")
}

func renderTemplate(writer http.ResponseWriter, name string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + name + ".tmpl")

	err := parsedTemplate.Execute(writer, nil)

	if err != nil {
		log.Println("Error parsing page:", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	log.Printf(fmt.Sprintf("Starting server on port %s...\n", portNumber))
	http.ListenAndServe(portNumber, nil)
}
