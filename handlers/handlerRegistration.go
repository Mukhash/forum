package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			if err := newUser.IsValid(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			newUser.Password = string(hashedPass)

			if err := db.InsertUser(env.db, &newUser); err != nil {
				http.Error(w, insertUserErrorText, http.StatusBadRequest)
			}

			cookie := utils.CreateCookie()
			if err := db.InsertCookie(env.db, cookie, newUser.ID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}

	})
}
