package handlers

import (
	"forum/utils"
	"net/http"
)

func (env *env) LogHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/login/" {
		http.NotFound(w, r)
		return
	}

	utils.RenderTemplate(w, env.tmpl, "login", nil)
}
