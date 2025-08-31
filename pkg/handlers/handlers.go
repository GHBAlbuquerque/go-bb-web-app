package handlers

import (
	"net/http"

	"github.com/GHBAlbuquerque/go-bb-web-app/pkg/config"
	"github.com/GHBAlbuquerque/go-bb-web-app/pkg/models"
	"github.com/GHBAlbuquerque/go-bb-web-app/pkg/render"
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

// NewHandlers -> Constructor setting Repo
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP) // everytime somebody hits the homepage, i'm storing the ip

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := map[string]string{}

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) GeneralsQuarters(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MarjorsSuite(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeReservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "search-availability.page.tmpl", &models.TemplateData{})
}
