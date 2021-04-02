package main

import (
	"cpratap/samplewebapp/pkg/config"
	"cpratap/samplewebapp/pkg/handlers"
	"cpratap/samplewebapp/pkg/render"
	"log"
	"net/http"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomePageHandler)
	_ = http.ListenAndServe(":8080", nil)
}
