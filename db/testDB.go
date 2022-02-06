package db

import (
	"database/sql"
	"forum/models"
	"math/rand"
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

	comments := []string{
		"great job!",
		"yeahh boiii",
		"you doing good job",
		"i agree",
		"i disagree",
		"you suck",
		"good morning",
		"good evening",
		"good night",
		"pleaaaase",
		"nooooooooooooo",
	}

	for _, v := range comments {
		comment := models.Comment{
			PostID:   rand.Intn(3) + 1,
			UserID:   rand.Intn(2) + 1,
			Username: "Test Username",
			Body:     v,
			Datefrom: time.Now(),
		}
		err := InsertComment(db, &comment)
		if err != nil {
			return err
		}
	}
	return nil
}
