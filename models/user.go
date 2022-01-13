package models

import "net/http"

type User struct {
	ID            int
	Email         string
	Name          string
	Password      string
	Authenticated bool
}

func (u *User) InitUser(r *http.Request) {
	u.Email = r.FormValue("email")
	u.Name = r.FormValue("username")
	u.Password = r.FormValue("password")
}
