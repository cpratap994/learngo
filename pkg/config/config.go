package config

import "html/template"

// holds application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}
