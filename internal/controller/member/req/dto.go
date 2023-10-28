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
	MemberId int  `json:"member_id"`
	IsBanned bool `json:"is_banned"`
}

func (d PatchMemberDto) ToDomain(userId, cafeId int) domain.Member {
	return domain.Member{
		Id:       d.MemberId,
		CafeId:   cafeId,
		UserId:   userId,
		IsBanned: d.IsBanned,
	}
}
