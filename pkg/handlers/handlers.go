package handlers

import (
	"cpratap/samplewebapp/pkg/config"
	"cpratap/samplewebapp/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) HomePageHandler(responseWriter http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(responseWriter, "home.page.tmpl")

}
