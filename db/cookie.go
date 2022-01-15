package db

import (
	"database/sql"
	"net/http"
)

func InsertCookie(db *sql.DB, cookie *http.Cookie, user_id int) error {
	_, err := db.Exec(
		"INSERT INTO auth_sessions (user_id, cookie_value, status, dateto) VALUES (?,?,?,?)",
		user_id, cookie.Value, 1, cookie.Expires,
	)
	return err
}
