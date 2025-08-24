package handlers

import (
	"go-bb-web-app/pkg/render"
	"net/http"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "home.page")
}

func About(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "about.page")
}
