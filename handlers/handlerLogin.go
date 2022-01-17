package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
)

func (env *env) LogHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(ctxUserKey).(*models.User)
		switch r.Method {
		case http.MethodGet:
			utils.RenderTemplate(w, env.tmpl, "login", user)
		case http.MethodPost:
			if user.Authenticated {
				http.Redirect(w, r, "/single_sign_on", http.StatusFound)
				return
			}
			pass := r.PostFormValue("password")
			login := r.PostFormValue("username")

			if pass == "" || login == "" {
				http.Error(w, "No pass or login", http.StatusBadRequest)
			}

			user, err := db.FindUserByEmail(env.db, login, pass)
			if err != nil {
				http.Error(w, "Incorrect username or password", http.StatusUnauthorized)
			}

			cookie := utils.CreateCookie()
			if err = db.InsertCookie(env.db, cookie, user.ID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
}
