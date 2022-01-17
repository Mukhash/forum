package models

import "time"

type Hashtag struct {
	ID       int
	Title    string
	Datefrom time.Time
}
