package handlers

import "net/http"

func (env *env) MainHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}
