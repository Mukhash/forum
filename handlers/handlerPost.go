package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
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

			post, err := db.GetPost(env.db, postId)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			postpage := models.Postpage{
				User: user,
				Post: post,
			}

			utils.RenderTemplate(w, env.tmpl, "post", postpage)

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
