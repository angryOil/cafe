package dto

import (
	"cafe/internal/domain"
)

type PatchDto struct {
	MemberId int  `json:"member_id"`
	IsBanned bool `json:"is_banned"`
}

func ToPatchDto(d domain.Member) PatchDto {
	return PatchDto{
		MemberId: d.Id,
		IsBanned: d.IsBanned,
	}
}

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

func ToJoinMemberDto(d domain.Member) JoinMemberDto {
	return JoinMemberDto{Nickname: d.Nickname}
}
