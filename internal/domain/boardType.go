package domain

import "time"

type BoardType struct {
	Id          int
	CreateBy    int
	CafeId      int
	Name        string
	Description string
	CreatedAt   time.Time
}
