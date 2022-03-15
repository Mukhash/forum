package utils

import (
	"forum/models"
	"net/http"
	"text/template"
)

func Error(w http.ResponseWriter, tmpl *template.Template, user *models.User, errCode int) {
	w.WriteHeader(errCode)
	errPage := models.Err{Text: http.StatusText(errCode), User: user}
	RenderTemplate(w, tmpl, "error", errPage)
}
