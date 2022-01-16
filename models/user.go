package models

import (
	"errors"
	"net/http"
	"regexp"
)

type User struct {
	ID            int
	Email         string
	Name          string
	Password      string
	Authenticated bool
}

func (u *User) InitUser(r *http.Request) error {
	if r.FormValue("email") == "" {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	u.Email = r.FormValue("email")

	if r.FormValue("username") == "" {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	u.Name = r.FormValue("username")

	if r.FormValue("password") == "" {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	u.Password = r.FormValue("password")
	return nil
}

func (u *User) IsValid() error {
	regMail := regexp.MustCompile(`^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`)
	if !regMail.MatchString(u.Email) {
		return errors.New("Invalid email")
	}
	if len(u.Password) < 6 {
		return errors.New("Too short password")
	}
	return nil
}
