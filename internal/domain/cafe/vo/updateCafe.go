package vo

import "time"

type UpdateCafe struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
	CreatedAt   time.Time
}
