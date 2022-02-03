package handlers

import (
	"forum/db"
	"forum/models"
	"net/http"
	"strconv"
	"time"
)

func (env *env) CommentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		if !user.Authenticated {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		switch r.Method {
		case http.MethodPost:
			body, postIDraw := r.PostFormValue("body"), r.PostFormValue("post_id")
			if body == "" || postIDraw == "" {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			postID, err := strconv.Atoi(postIDraw)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			comment := &models.Comment{
				PostID:   postID,
				UserID:   user.ID,
				Username: user.Name,
				Body:     body,
				Datefrom: time.Now(),
			}

			if err := db.InsertComment(env.db, comment); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

		}
	})
}
