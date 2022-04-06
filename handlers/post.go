package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (env *env) PostHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/post/" {
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}

			postId, err := strconv.Atoi(r.URL.Path[6:])
			if postId <= 0 || err != nil {
				utils.Error(w, env.tmpl, user, http.StatusNotFound)
				return
			}

			post, err := db.GetPost(env.db, postId)
			if err != nil {
				utils.Error(w, env.tmpl, user, http.StatusNotFound)
				return
			}

			postpage := models.Postpage{
				User: user,
				Post: post,
			}

			utils.RenderTemplate(w, env.tmpl, "post", postpage)

		case http.MethodPost:
			if !user.Authenticated {
				utils.Error(w, env.tmpl, user, http.StatusUnauthorized)
				return
			}

			post := &models.Post{UserID: user.ID, Username: user.Name}
			body := r.PostFormValue("body")
			if body == "" || strings.TrimSpace(body) == "" {
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}
			post.Body = body
			post.Datefrom = time.Now()

			err := db.CreatePost(env.db, post)
			if err != nil {
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		default:
			utils.Error(w, env.tmpl, user, http.StatusMethodNotAllowed)
			return
		}
	})
}
