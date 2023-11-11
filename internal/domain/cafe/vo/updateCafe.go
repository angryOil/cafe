package vo

import "time"

type UpdateCafe struct {
	Id          int       // 카페 식별자
	OwnerId     int       // 카페 주인 식별자
	Name        string    // 카페 이름
	Description string    // 카페 설명
	CreatedAt   time.Time // 카페 개설된 시간
}
