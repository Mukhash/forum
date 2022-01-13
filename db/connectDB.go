package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectBD() (*sql.DB, error) {
	os.Remove("forum.db")
	file, err := os.Create("forum.db")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	forumDB, err := sql.Open("sqlite3", "./forum.db")
	if err != nil {
		return nil, err
	}

	for _, query := range getQuery() {
		_, err := forumDB.Exec(query)
		if err != nil {
			return nil, err
		}
	}

	return forumDB, nil
}
