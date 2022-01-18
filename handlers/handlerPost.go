package handlers

import (
	"forum/db"
	"forum/models"
	"net/http"
)

func (env *env) PostHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		case http.MethodPost:
			user := r.Context().Value(ctxUserKey).(*models.User)

			if !user.Authenticated {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			post := &models.Post{UserID: user.ID, Username: user.Name}
			body := r.PostFormValue("body")
			if body == "" {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			}
			post.Body = body

			_, err := db.CreatePost(env.db, post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		default:
		}
	})
}
