package handlers

import (
	"database/sql"
	"text/template"
)

type env struct {
	db   *sql.DB
	tmpl *template.Template
}

func InitEnv() *env {
	return &env{}
}

func (env *env) SetDB(db *sql.DB) {
	env.db = db
}

func (env *env) SetTmpl(tmpl *template.Template) {
	env.tmpl = tmpl
}
