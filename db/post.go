package db

import (
	"database/sql"
	"forum/models"
)

func GetPosts(db *sql.DB) (*[]models.Post, error) {
	posts := &[]models.Post{}

	query := "select id, user_id, text, datefrom from posts order by datefrom DESC limit 1000"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.UserID, &post.Body, &post.Datefrom)
		if err != nil {
			return nil, err
		}
		*posts = append(*posts, post)
	}
	return posts, nil
}

func CreatePost(db *sql.DB, post *models.Post) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec(
		"INSERT INTO posts (user_id, body, creation_date) VALUES(?, ?, ?)",
		post.UserID, post.Body, post.Datefrom,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
