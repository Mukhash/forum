package models

import "time"

type Post struct {
	ID            int64
	UserID        int
	Username      string
	Body          string
	Datefrom      time.Time
	LikeType      int
	LikesCount    int
	CommentsCount int
	Tag           Tag
}

type PostFeed struct {
	Posts       *[]Post `json:"data"`
	NextFirstId int64   `json:"nextFirstId"`
}
