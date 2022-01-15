package handlers

import (
	"context"
	"forum/db"
	"forum/models"
	"net/http"
)

type ctxKey int

const ctxUserKey ctxKey = iota
const cookieName = "auth_session"

func (env *env) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{Name: "Guest"}
		cookie, err := r.Cookie(cookieName)
		if err != http.ErrNoCookie {
			user, _ = db.FindUserBySession(env.db, cookie.Value)
		}

		r2 := r.Clone(context.WithValue(r.Context(), ctxUserKey, user))
		next.ServeHTTP(w, r2)
	})
}
