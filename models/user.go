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
	formVal := r.PostFormValue("email")
	if formVal == "" {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	u.Email = formVal

	if formVal = r.PostFormValue("username"); formVal == "" {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	u.Name = formVal

	if formVal = r.PostFormValue("password"); formVal == "" {
		return errors.New(http.StatusText(http.StatusBadRequest))
	}
	u.Password = formVal
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
