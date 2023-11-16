package vo

import "time"

type Update struct {
	Id            int
	BoardType     int
	CafeId        int
	Writer        int
	Title         string
	Content       string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
