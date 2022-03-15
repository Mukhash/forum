package handlers

import (
	"fmt"
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (env *env) CommentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		if !user.Authenticated {
			utils.Error(w, env.tmpl, user, http.StatusUnauthorized)
			return
		}
		switch r.Method {
		case http.MethodPost:
			body, postIDraw := r.PostFormValue("body"), r.PostFormValue("post_id")
			if body == "" || postIDraw == "" || strings.TrimSpace(body) == "" {
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}
			postID, err := strconv.Atoi(postIDraw)
			if err != nil {
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}
			fmt.Println(postID, user.ID)
			comment := &models.Comment{
				PostID:   postID,
				UserID:   user.ID,
				Username: user.Name,
				Body:     body,
				Datefrom: time.Now(),
			}

			if err := db.InsertComment(env.db, comment); err != nil {
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}

			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
		case http.MethodGet:
			utils.Error(w, env.tmpl, user, http.StatusMethodNotAllowed)
			return
		}
	})
}
