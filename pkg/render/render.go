package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(responseWriter http.ResponseWriter, tmpl string) {

	baseMap, err := renderBaseLayout(responseWriter)
	if err != nil {
		fmt.Println("Error occured")
		return
	}
	fmt.Println(baseMap)

	parsedTemplate, _ := template.ParseFiles("../../templates/" + tmpl)
	err = parsedTemplate.Execute(responseWriter, nil)

	if err != nil {
		fmt.Println("Error occured while parsing template", err)
	}
}

func renderBaseLayout(w http.ResponseWriter) (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		fmt.Println(name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
