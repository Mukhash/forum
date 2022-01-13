package utils

import (
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, tmplName string, obj interface{}) {
	err := tmpl.ExecuteTemplate(w, tmplName+".html", obj)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError)+" "+err.Error(), http.StatusInternalServerError)
		return
	}
}
