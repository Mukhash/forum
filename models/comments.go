package models

import "time"

type Commment struct {
	ID       int
	PostID   int
	UserID   int
	Username string
	Body     string
	Datefrom time.Time
}
