package db

import (
	"database/sql"
	"forum/models"
	"time"
)

// FillDatabase with example data
func FillDatabase(db *sql.DB) error {
	users := [][]string{
		{
			"talgat.mukhash@gmail.com",
			"Talgat",
			"qwerty",
		},

		{
			"muhash.astana@mail.ru",
			"Muhash",
			"qwerty",
		},
	}

	for _, v := range users {
		user := &models.User{
			Email:    v[0],
			Name:     v[1],
			Password: v[2],
		}
		err := InsertUser(db, user)
		if err != nil {
			return err
		}
	}

	texts := []string{
		"lorem ipsum",
		"Dossan mal",
		"Curiosity killed the cat",
		"Поехали",
		"L is Kira",
		"Fuck Alem",
		"I want to work in Moscow",
		"Let's go Ozon",
		"Let's go Technodom",
	}

	for _, v := range texts {
		post := &models.Post{
			UserID:        1,
			Body:          v,
			Datefrom:      time.Now(),
			CommentsCount: 15,
		}
		err := CreatePost(db, post)
		if err != nil {
			return err
		}

	}
	return nil
}
