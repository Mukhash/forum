package handlers

import (
	"database/sql"
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"net/url"
)

func (env *env) SearchHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		if !user.Authenticated {
			utils.Error(w, env.tmpl, user, http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodGet {
			utils.Error(w, env.tmpl, user, http.StatusMethodNotAllowed)
			return
		}

		var posts []models.Post
		var err error
		queries, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}

		switch true {
		case queries.Get("data-choice") != "":
			posts, err = getTags(w, r, queries.Get("data-choice"), env.db)
			if err != nil {
				utils.Error(w, env.tmpl, user, http.StatusInternalServerError)
				return
			}
		case queries.Get("favs") == "favs":
			posts, err = getFavs(w, r, user.ID, env.db)
			if err != nil {
				utils.Error(w, env.tmpl, user, http.StatusInternalServerError)
				return
			}
		case queries.Get("created") == "created":
			posts, err = getCreated(w, r, user.ID, env.db)
			if err != nil {
				utils.Error(w, env.tmpl, user, http.StatusInternalServerError)
				return
			}
		default:
			utils.Error(w, env.tmpl, user, http.StatusBadRequest)
			return
		}

		page := models.Mainpage{User: user, Posts: &posts}
		utils.RenderTemplate(w, env.tmpl, "index_search", page)
	})
}

func getTags(w http.ResponseWriter, r *http.Request, query string, mdb *sql.DB) ([]models.Post, error) {
	tagID, err := db.GetTagByTitle(mdb, query)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	if tagID == 0 {
		return posts, nil
	}
	posts, err = db.GetPostsByTag(mdb, tagID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func getFavs(w http.ResponseWriter, r *http.Request, userID int, mdb *sql.DB) ([]models.Post, error) {
	posts, err := db.GetFavs(mdb, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func getCreated(w http.ResponseWriter, r *http.Request, userID int, mdb *sql.DB) ([]models.Post, error) {
	posts, err := db.GetCreated(mdb, userID)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
