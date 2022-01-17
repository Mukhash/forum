package models

import "time"

type Post struct {
	ID            int
	UserID        int
	Username      string
	Body          string
	Datefrom      time.Time
	LikesCount    int
	CommentsCount int
}
