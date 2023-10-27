package req

import (
	"cafe/internal/domain"
)

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

func (d JoinMemberDto) ToDomain(userId, cafeId int) domain.Member {
	return domain.Member{
		CafeId:   cafeId,
		UserId:   userId,
		Nickname: d.Nickname,
	}
}
