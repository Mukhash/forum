package handlers

import (
	"context"
	"forum/models"
	"net/http"
)

type ctxKey int

const ctxUserKey ctxKey = iota

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := models.User{Name: "Guest"}
		_, err := r.Cookie("authSession")
		if err != http.ErrNoCookie {
			//user = db.FindUserBySession
		}

		r2 := r.Clone(context.WithValue(r.Context(), ctxUserKey, user))
		next.ServeHTTP(w, r2)
	})
}
