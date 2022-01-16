package db

import (
	"database/sql"
	"forum/models"

	"golang.org/x/crypto/bcrypt"
)

// InsertUser inserts newly created user into users table in forum db.
func InsertUser(db *sql.DB, user *models.User) error {
	var password string
	res, err := db.Exec(
		"INSERT INTO users (name,email,password) VALUES(?,?,?)",
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

func FindUserBySession(db *sql.DB, cookieValue string) (*models.User, error) {
	query :=
		`select users.id, users.name
	from users
	left join auth_sessions as auth
	on users.id = auth.user_id
	where auth.cookie_value = ?`

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
	query := "select id, name, password from users where email = ?"
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return user, err
}
