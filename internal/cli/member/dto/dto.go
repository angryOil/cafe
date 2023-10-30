package dto

import (
	"cafe/internal/domain"
	"time"
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

type MemberInfoDto struct {
	Id        int    `json:"member_id,omitempty"`
	UserId    int    `json:"user_id"`
	NickName  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func ToMemberDomainList(dtos []MemberInfoDto) []domain.Member {
	results := make([]domain.Member, len(dtos))
	for i, d := range dtos {
		results[i] = domain.Member{
			Id:       d.Id,
			UserId:   d.UserId,
			Nickname: d.NickName,
		}
	}
	return results
}

func ToMemberInfoDtoList(domains []domain.Member) []MemberInfoDto {
	results := make([]MemberInfoDto, len(domains))
	for i, d := range domains {
		results[i] = MemberInfoDto{
			Id:       d.Id,
			UserId:   d.UserId,
			NickName: d.Nickname,
		}
	}
	return results
}

func (dto MemberInfoDto) ToDomain() domain.Member {
	t, err := convertToTime(dto.CreatedAt)
	if err != nil {
		t = time.Time{}
	}
	return domain.Member{
		Id:        dto.Id,
		UserId:    dto.UserId,
		Nickname:  dto.NickName,
		CreatedAt: t,
	}
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertToTime(s string) (time.Time, error) {
	t, err := time.ParseInLocation(time.RFC3339, s, koreaZone)
	return t, err
}
