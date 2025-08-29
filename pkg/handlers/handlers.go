package handlers

import (
	"go-bb-web-app/pkg/config"
	"go-bb-web-app/pkg/models"
	"go-bb-web-app/pkg/render"
	"net/http"
)

// Repository is the type definition
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepository -> Constructor setting AppConfig
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP) // everytime somebody hits the homepage, i'm storing the ip

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello again!"

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

// NewHandlers -> Constructor setting Repo
func NewHandlers(r *Repository) {
	Repo = r
}
