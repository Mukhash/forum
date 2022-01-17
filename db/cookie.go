package db

import (
	"database/sql"
	"net/http"
)

// InsertCookie adds login session data into database
func InsertCookie(db *sql.DB, cookie *http.Cookie, user_id int) error {
	_, err := db.Exec(
		"INSERT INTO auth_sessions (user_id, cookie_value, status, dateto) VALUES (?,?,?,?)",
		user_id, cookie.Value, 1, cookie.Expires,
	)
	return err
}

// DeleteUserSession is called only when sign out
func DeleteUserSession(db *sql.DB, userID int) error {
	query := "delete from auth_sessions where user_id = ?"
	_, err := db.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
