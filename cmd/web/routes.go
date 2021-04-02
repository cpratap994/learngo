package main

import (
	"cpratap/samplewebapp/pkg/config"
	"cpratap/samplewebapp/pkg/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func routePath(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.HomePageHandler))
	return mux
}
