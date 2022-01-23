package handlers

import (
	"encoding/json"
	"fmt"
	"forum/db"
	"forum/models"
	"net/http"
	"strconv"
)

func (env *env) NextPostsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		last := (*posts)[len(*posts)-1].ID
		postsFeed := models.PostFeed{Posts: posts, NextFirstId: last - 1}
		postsJson, err := json.Marshal(postsFeed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", string(postsJson))
	})
}
