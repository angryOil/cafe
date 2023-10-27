package res

import (
	"cafe/internal/domain"
	"time"
)

type MemberInfoDto struct {
	Id        int    `json:"member_id,omitempty"`
	Nickname  string `json:"nick_name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	IsBanned  bool   `json:"is_banned,omitempty"`
}

func ToMemberInfoDto(d domain.Member) MemberInfoDto {
	return MemberInfoDto{
		Id:        d.Id,
		Nickname:  d.Nickname,
		CreatedAt: convertTimeToString(d.CreatedAt),
		IsBanned:  d.IsBanned,
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
