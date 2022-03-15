package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"strconv"
)

func (env *env) CommentLikeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		if !user.Authenticated {
			utils.Error(w, env.tmpl, user, http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodPost {
			utils.Error(w, env.tmpl, user, http.StatusMethodNotAllowed)
			return
		}

		commentIdRaw := r.PostFormValue("comment_id")
		if commentIdRaw == "" {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}
		commentID, err := strconv.Atoi(commentIdRaw)
		if err != nil {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}

		likeTypeRaw := r.PostFormValue("action")
		if likeTypeRaw == "" {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}
		likeType, err := strconv.Atoi(likeTypeRaw)
		if err != nil {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}

		likeComment := models.LikeComment{CommentID: commentID, UserID: user.ID, Type: likeType}

		if err = db.InsertLike(env.db, &likeComment); err != nil {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}

		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	})
}
