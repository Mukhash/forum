package handlers

import (
	"forum/utils"
	"net/http"
)

func (env *env) TestIndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.RenderTemplate(w, env.tmpl, "index_test", nil)
		return
	})
}
