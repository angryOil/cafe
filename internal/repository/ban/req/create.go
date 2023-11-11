package req

import "time"

type Create struct {
	UserId      int
	MemberId    int
	CafeId      int
	Description string
	CreatedAt   time.Time
}
