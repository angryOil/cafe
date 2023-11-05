package request

import "time"

type CreateCafe struct {
	OwnerId     int
	Name        string
	Description string
	CreatedAt   time.Time
}

type UpdateCafe struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
	CreatedAt   time.Time
}
