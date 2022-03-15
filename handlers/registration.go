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
			user := r.Context().Value(ctxUserKey).(*models.User)
			if user.Authenticated {
				http.Redirect(w, r, "/reg_sign_on", http.StatusFound)
				return
			}

			utils.RenderTemplate(w, env.tmpl, "register", nil)
		case http.MethodPost:
			newUser := models.User{}
			defer r.Body.Close()

			if err := newUser.InitUser(r); err != nil {
				utils.Error(w, env.tmpl, &newUser, http.StatusBadRequest)
				return
			}

			if err := newUser.IsValid(); err != nil {
				utils.Error(w, env.tmpl, &newUser, http.StatusBadRequest)
				return
			}

			if err := db.InsertUser(env.db, &newUser); err != nil {
				utils.Error(w, env.tmpl, &newUser, http.StatusInternalServerError)
				return
			}

			cookie := utils.CreateCookie()
			if err := db.InsertCookie(env.db, cookie, newUser.ID); err != nil {
				utils.Error(w, env.tmpl, &newUser, http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		default:
			utils.Error(w, env.tmpl, &models.User{Authenticated: false, Name: "Guest"}, http.StatusMethodNotAllowed)
			return
		}

	})
}
