package handlers

import (
	"encoding/json"
	"fmt"
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"strconv"
)

func (env *env) NextPostsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			utils.Error(w, env.tmpl, &models.User{}, http.StatusMethodNotAllowed)
			return
		}
		firstID, err := strconv.Atoi(r.URL.Query().Get("first_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		posts, err := db.GetNextPosts(env.db, firstID, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var last int64
		if len(*posts) != 0 {
			last = (*posts)[len(*posts)-1].ID
		}

		postsFeed := models.PostFeed{Posts: posts, NextFirstId: last - 1}
		postsJson, err := json.Marshal(postsFeed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = fmt.Fprintf(w, "%s", string(postsJson))
		if err != nil {
			return
		}
	})
}
