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

type PatchMemberDto struct {
	Nickname string `json:"nickname"`
}

func (d PatchMemberDto) ToDomain(userId, cafeId int) domain.Member {
	return domain.Member{
		Nickname: d.Nickname,
		CafeId:   cafeId,
		UserId:   userId,
	}
}
