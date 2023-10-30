package domain

import "time"

type Ban struct {
	Id          int
	UserId      int
	MemberId    int
	CafeId      int
	Description string
	CreatedAt   time.Time
}
