package dto

import (
	"cafe/internal/domain"
)

type PatchDto struct {
	Nickname string `json:"nickname"`
}

func ToPatchDto(d domain.Member) PatchDto {
	return PatchDto{
		Nickname: d.Nickname,
	}
}

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

func ToJoinMemberDto(d domain.Member) JoinMemberDto {
	return JoinMemberDto{Nickname: d.Nickname}
}
