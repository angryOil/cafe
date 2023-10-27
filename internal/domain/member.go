package domain

import "time"

type Member struct {
	Id        int       `json:"member_id,omitempty"`
	CafeId    int       `json:"cafe_id,omitempty"`
	UserId    int       `json:"user_id,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	IsBanned  bool      `json:"is_banned,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
