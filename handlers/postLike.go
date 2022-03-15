package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"strconv"
)

func (env *env) PostLikeHandler() http.Handler {
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

		// Querying post_id
		postIdRaw := r.PostFormValue("post_id")
		if postIdRaw == "" {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}
		postID, err := strconv.Atoi(postIdRaw)
		if err != nil {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}
		// Querying action
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
		// Creating models.LikePost object
		likePost := models.LikePost{PostID: postID, UserID: user.ID, Type: likeType}
		if err = db.InsertLike(env.db, &likePost); err != nil {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}
		// Redirecting to Referer
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusFound)
	})
}
