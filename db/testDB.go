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
			"tm@gmail.com",
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
		"lorem ipsum #1",
		"Alem Cup",
		"Curiosity killed the cat",
		"Поехали",
		"L is Kira",
		"Hello World",
		"I want to work in Nasa",
		"Let's go GO",
		"Let's go Ruby",
	}

	for _, v := range texts {
		post := &models.Post{
			UserID:   rand.Intn(2) + 1,
			Body:     v,
			Datefrom: time.Now(),
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
