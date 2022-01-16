package handlers

import (
	"forum/db"
	"forum/models"
	"forum/utils"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

			if err := newUser.InitUser(r); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			if err := newUser.IsValid(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			hashedPass, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
			newUser.Password = string(hashedPass)

			if err := db.InsertUser(env.db, &newUser); err != nil {
				http.Error(w, insertUserErrorText, http.StatusBadRequest)
			}

			uuid := uuid.NewV4()
			dateto := time.Now().Add(time.Duration(7 * 24 * time.Hour))
			cookie := http.Cookie{Name: cookieName, Value: uuid.String(), Expires: dateto}
			if err := db.InsertCookie(env.db, &cookie, newUser.ID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			http.SetCookie(w, &cookie)

			http.Redirect(w, r, "/", http.StatusFound)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}

	})
}
