package models

import "time"

type Post struct {
	ID       int
	UserID   int
	Body     string
	Datefrom time.Time
}
