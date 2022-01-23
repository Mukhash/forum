package models

import "time"

type Post struct {
	ID            int64
	UserID        int
	Username      string
	Body          string
	Datefrom      time.Time
	LikesCount    int
	CommentsCount int
}

type PostFeed struct {
	Posts       *[]Post `json:"data"`
	NextFirstId int64   `json:"nextFirstId"`
}
