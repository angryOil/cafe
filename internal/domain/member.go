package domain

import "time"

type Member struct {
	Id        int       `json:"member_id,omitempty"`
	CafeId    int       `json:"cafe_id,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type MemberListCount struct {
	Members []Member `json:"members"`
	Count   int      `json:"count"`
}
