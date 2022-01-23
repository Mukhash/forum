package utils

import "text/template"

func GetTmpl() (*template.Template, error) {
	tmpl, err := template.ParseGlob("./static/templates/*.html")
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
