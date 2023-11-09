package vo

import "time"

type Info struct {
	Id          int
	UserId      int
	MemberId    int
	CafeId      int
	Description string
	CreatedAt   time.Time
}
