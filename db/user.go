package db

import (
	"database/sql"
	"forum/models"
)

// InsertUser inserts newly created user into users table in forum db.
func InsertUser(db *sql.DB, user *models.User) error {
	var password string
	res, err := db.Exec(
		"INSERT INTO users (username,email,password) VALUES(?,?,?)",
		user.Name, user.Email, string(password),
	)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	user.ID = int(id)
	return err
}

// DeleteUserByName deletes user by name from users table in forum db.
func DeleteUserByName(db *sql.DB, name string) error {
	_, err := db.Exec("DELETE FROM users WHERE name=?", name)
	return err
}

// DeleteUserByName deletes user by id from users table in forum db.
func DeleteUserByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
