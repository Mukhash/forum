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
				utils.Error(w, env.tmpl, user, http.StatusBadRequest)
				return
			}

			user, err := db.FindUserByEmail(env.db, login, pass)
			if err != nil {
				user = &models.User{Name: "Guest"}
				utils.Error(w, env.tmpl, user, http.StatusUnauthorized)
				return
			}

			cookie := utils.CreateCookie()
			if err = db.DeleteUserSession(env.db, user.ID); err != nil {
				utils.Error(w, env.tmpl, user, http.StatusInternalServerError)
				return
			}
			if err = db.InsertCookie(env.db, cookie, user.ID); err != nil {
				utils.Error(w, env.tmpl, user, http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		default:
			utils.Error(w, env.tmpl, user, http.StatusMethodNotAllowed)
			return
		}
	})
}
