package handlers

import (
	"forum/utils"
	"net/http"
)

func (env *env) HandleRegSignOn(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, env.tmpl, "reg_sign_on", nil)
}
