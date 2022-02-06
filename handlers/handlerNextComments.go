package handlers

import (
	"encoding/json"
	"fmt"
	"forum/db"
	"forum/models"
	"net/http"
	"strconv"
)

func (env *env) NextComments() http.Handler {
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

		postID, err := strconv.Atoi(r.URL.Query().Get("post_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		comments, err := db.GetComments(env.db, postID, firstID, limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		last := (*comments)[len(*comments)-1].ID
		commentsFeed := models.CommentFeed{Comments: comments, NextFirstId: int64(last)}
		commentsJson, err := json.Marshal(commentsFeed)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", string(commentsJson))
	})
}
