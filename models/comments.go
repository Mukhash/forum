package models

import "time"

type Comment struct {
	ID         int
	PostID     int
	UserID     int
	Username   string
	Body       string
	Datefrom   time.Time
	LikesCount int
}

type CommentFeed struct {
	Comments    *[]Comment `json:"data"`
	NextFirstId int64      `json:"nextFirstId"`
}
