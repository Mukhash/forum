package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
)

func (env *env) MainHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		if !user.Authenticated {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		posts, err := db.GetPosts(env.db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		mainpage := models.Mainpage{User: user, Posts: posts}
		utils.RenderTemplate(w, env.tmpl, "index", mainpage)
	})
}
