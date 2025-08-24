package main

import (
	"log"
	"net/http"
	"text/template"
)

func renderTemplate(writer http.ResponseWriter, name string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + name + ".tmpl")

	err := parsedTemplate.Execute(writer, nil)

	if err != nil {
		log.Println("Error parsing page:", err)
	}
}
