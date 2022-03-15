package handlers

import (
	"context"
	"fmt"
	"forum/db"
	"forum/models"
	"forum/utils"
	"log"
	"net/http"
	"time"
)

type ctxKey int

const ctxUserKey ctxKey = iota

func (env *env) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		request := fmt.Sprintf("%s %s", r.Method, r.URL)

		user := &models.User{Name: "Guest"}
		cookie, err := r.Cookie(utils.CookieName)
		if err != http.ErrNoCookie {
			user, _ = db.FindUserBySession(env.db, cookie.Value)
		}

		r2 := r.Clone(context.WithValue(r.Context(), ctxUserKey, user))
		next.ServeHTTP(w, r2)
		log.Println(request, ": success", time.Since(now))
	})
}
