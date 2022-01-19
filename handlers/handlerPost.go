package handlers

import (
	"forum/db"
	"forum/models"
	"net/http"
)

func (env *env) PostHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		switch r.Method {
		case http.MethodGet:

			if !user.Authenticated {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// page := r.URL.Query().Get("page")
			// limit := r.URL.Query().Get("limit")

		case http.MethodPost:

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

			err := db.CreatePost(env.db, post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Redirect(w, r, "/", http.StatusFound)
		default:
		}
	})
}
