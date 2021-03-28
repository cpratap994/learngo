package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(responseWriter http.ResponseWriter, tmpl string) {

	parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)
	err := parsedTemplate.Execute(responseWriter, nil)

	if err != nil {
		fmt.Println("Error occured while parsing template", err)
	}
}
