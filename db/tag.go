package db

import "database/sql"

func GetTagByTitle(db *sql.DB, title string) (int, error) {
	query := `SELECT id FROM tags
	WHERE title = ?`

	row := db.QueryRow(query, title)

	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return -1, err
	}
	return id, nil
}
