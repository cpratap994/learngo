package main

import (
	"cpratap/samplewebapp/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomePageHandler)
	_ = http.ListenAndServe(":8080", nil)
}
