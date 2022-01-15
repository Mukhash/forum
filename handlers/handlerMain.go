package handlers

import (
	"forum/models"
	"forum/utils"
	"net/http"
)

func (env *env) MainHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		mainpage := models.Mainpage{User: user}
		utils.RenderTemplate(w, env.tmpl, "index", mainpage)
	})
}
