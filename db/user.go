package db

import (
	"database/sql"
	"forum/models"

	"golang.org/x/crypto/bcrypt"
)

// InsertUser inserts newly created user into users table in forum db.
func InsertUser(db *sql.DB, user *models.User) error {
	cryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	res, err := db.Exec(
		"INSERT INTO users (name,email,password) VALUES(?,?,?)",
		user.Name, user.Email, string(cryptedPassword),
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

func FindUserBySession(db *sql.DB, cookieValue string) (*models.User, error) {
	query :=
		`SELECT users.id, users.name
	FROM users
	LEFT JOIN auth_sessions AS auth
	ON users.id = auth.user_id
	WHERE auth.cookie_value = ?`

	user := models.User{Authenticated: true}

	err := db.QueryRow(query, cookieValue).Scan(&user.ID, &user.Name)
	if err != nil {
		user.Authenticated = false
		user.Name = "Guest"
		if err == sql.ErrNoRows {
			err = nil
		}
	}
	return &user, err
}

func FindUserByEmail(db *sql.DB, email string, password string) (*models.User, error) {
	user := &models.User{}
	var hashedPass string
	query := "SELECT id, name, password FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &hashedPass)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return user, err
}
