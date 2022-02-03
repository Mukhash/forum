package handlers

import (
	"forum/db"
	"forum/models"
	"net/http"
	"strconv"
	"time"
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
			if r.URL.Path == "/post/" {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			postId, err := strconv.Atoi(r.URL.Path[6:])
			if postId <= 0 || err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

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
			post.Datefrom = time.Now()

			err := db.CreatePost(env.db, post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			http.Redirect(w, r, "/", http.StatusFound)
		default:
		}
	})
}
