package db

import (
	"database/sql"
	"fmt"
	"forum/models"
)

func GetComments(db *sql.DB, post_id, firstID, limit int) (*[]models.Comment, error) {

	comments := []models.Comment{}
	queryWhere :=
		`SELECT id, post_id, user_id, username, body, datefrom,
		(
			SELECT IFNULL(
				SUM(CASE WHEN like_type = 1 THEN 1
					WHEN like_type = 2 then -1 ELSE 0
					END), 0)
			FROM  comment_likes as cl
			WHERE cl.comment_id = comments.id
		) as likes_count
		FROM comments
		WHERE id <= ? and post_id = ?
		ORDER BY datefrom DESC
		LIMIT ?`
	queryAll :=
		`SELECT *,
		(
			SELECT IFNULL(
				SUM(CASE WHEN like_type = 1 THEN 1
					WHEN like_type = 2 then -1 ELSE 0
					END), 0)
			FROM  comment_likes as cl
			WHERE cl.comment_id = comments.id
		) as likes_count
		FROM comments
		WHERE post_id = ?
		ORDER BY datefrom DESC
		LIMIT ?`
	query := queryWhere
	args := []interface{}{firstID, post_id, limit}

	if firstID == -1 {
		query = queryAll
		args = []interface{}{post_id, limit}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("error closing row: comment.go")
		}
	}(rows)
	count := 0
	for rows.Next() {
		count++
		comment := models.Comment{}
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Username, &comment.Body, &comment.Datefrom, &comment.LikesCount)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}
	return &comments, nil
}

func InsertComment(db *sql.DB, comment *models.Comment) error {
	_, err := db.Exec(
		`INSERT INTO comments 
		(post_id, user_id, username, body, datefrom)
		VALUES(?,?,?,?,?)`,
		comment.PostID, comment.UserID, comment.Username, comment.Body, comment.Datefrom)
	return err
}
