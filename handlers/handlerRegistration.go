package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
)

const insertUserErrorText = "Username or email already exists"

func (env *env) RegHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			utils.RenderTemplate(w, env.tmpl, "register", nil)
		case http.MethodPost:
			newUser := models.User{}
			defer r.Body.Close()
			newUser.InitUser(r)

			if err := db.InsertUser(env.db, &newUser); err != nil {
				http.Error(w, insertUserErrorText, http.StatusBadRequest)
			}

			http.Redirect(w, r, "/", http.StatusFound)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}

	})
}
