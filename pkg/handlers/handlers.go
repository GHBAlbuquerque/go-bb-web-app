package handlers

import (
	"go-bb-web-app/pkg/config"
	"go-bb-web-app/pkg/render"
	"net/http"
)

// Repository -> type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "home.page.tmpl")
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "about.page.tmpl")
}

// NewRepository -> Constructor setting AppConfig
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers -> Constructor setting Repo
func NewHandlers(r *Repository) {
	Repo = r
}
