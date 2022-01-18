package handlers

import (
	"forum/utils"
	"net/http"
)

func (env *env) TestHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, env.tmpl, "recycler", nil)
	return
}
