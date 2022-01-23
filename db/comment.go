package db

import (
	"database/sql"
	"forum/models"
)

func GetComments(db *sql.DB, post_id, firstID, limit int) (*[]models.Commment, error) {

	comments := []models.Commment{}
	queryWhere :=
		`SELECT id, post_id, user_id, body, datefrom FROM comments
		WHERE id <= ?
		ORDER BY datefrom DESC
		LIMIT ?`
	queryAll :=
		`SELECT * FROM comments
		ORDER BY datefrom DESC
		LIMIT ?`
	query := queryWhere
	args := []interface{}{firstID, limit}

	if firstID == -1 {
		query = queryAll
		args = []interface{}{1}
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := models.Commment{}
		err := rows.Scan(&comment.ID, &comment.UserID, &comment.Body, &comment.Datefrom)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}
	return &comments, nil
}
