package render

import "html/template"

func LoadTemplates() (*template.Template, error) {
	return template.ParseGlob("web/templates/pages/*.html")
}