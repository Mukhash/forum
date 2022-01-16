package utils

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

const CookieName = "auth_cookie"

func CreateCookie() *http.Cookie {
	uuid := uuid.NewV4()
	dateto := time.Now().Add(time.Duration(7 * 24 * time.Hour))
	cookie := &http.Cookie{Name: CookieName, Value: uuid.String(), Expires: dateto}
	return cookie
}
