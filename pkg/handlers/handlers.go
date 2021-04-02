package handlers

import (
	"cpratap/samplewebapp/pkg/render"
	"net/http"
)

func HomePageHandler(responseWriter http.ResponseWriter, request *http.Request) {

	render.RenderTemplate(responseWriter, "home.page.tmpl")

}
