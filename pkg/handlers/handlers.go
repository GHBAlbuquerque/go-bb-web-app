package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

	render.RenderTemplate(writer, request, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := map[string]string{}

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, request, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) GeneralsQuarters(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MarjorsSuite(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) MakeReservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "make-reservation.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "search-availability.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostAvailability(writer http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")
	adults := request.Form.Get("adults")
	children := request.Form.Get("children")
	room := request.Form.Get("room")

	text := fmt.Sprintf("start date is %s, end date is %s, adults: %s, children: %s, room is %s", start, end, adults, children, room)
	log.Println(text)

	resp := jsonResponse{
		OK:      true,
		Message: text,
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}
