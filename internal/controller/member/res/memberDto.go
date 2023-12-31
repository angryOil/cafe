package res

import (
	"cafe/internal/domain"
	"time"
)

type MemberInfoDto struct {
	Id        int    `json:"member_id,omitempty"`
	UserId    int    `json:"user_id"`
	Nickname  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func ToMemberInfoDto(d domain.Member) MemberInfoDto {
	return MemberInfoDto{
		Id:        d.Id,
		UserId:    d.UserId,
		Nickname:  d.Nickname,
		CreatedAt: convertTimeToString(d.CreatedAt),
	}
}

func ToMemberInfoDtoList(domains []domain.Member) []MemberInfoDto {
	results := make([]MemberInfoDto, len(domains))
	for i, d := range domains {
		results[i] = ToMemberInfoDto(d)
	}
	return results
}

type IdsTotalDto struct {
	Ids   []int `json:"ids"`
	Total int   `json:"total"`
}

func ToIdsTotalDto(d domain.IdsTotalDomain) IdsTotalDto {
	return IdsTotalDto{
		Ids:   d.Ids,
		Total: d.Total,
	}
}

func NewIdsTotalDto(ids []int, total int) IdsTotalDto {
	return IdsTotalDto{
		Ids:   ids,
		Total: total,
	}
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}
