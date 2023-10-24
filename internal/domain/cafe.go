package domain

import "time"

type Cafe struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
	CreatedAt   time.Time
}
