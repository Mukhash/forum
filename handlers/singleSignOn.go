package handlers

import (
	"forum/utils"
	"net/http"
)

func (env *env) HandleSignOn(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, env.tmpl, "single_sign_on", nil)
}
