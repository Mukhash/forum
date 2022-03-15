package handlers

import (
	"forum/db"
	"forum/models"
	"net/http"
)

func (env *env) LogoutHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		user := r.Context().Value(ctxUserKey).(*models.User)
		if !user.Authenticated {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		err := db.DeleteUserSession(env.db, user.ID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})
}
