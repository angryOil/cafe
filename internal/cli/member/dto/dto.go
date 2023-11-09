package dto

import (
	"cafe/internal/cli/member/req"
	"time"
)

type PatchDto struct {
	Nickname string `json:"nickname"`
}

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

func ToJoinMemberDto(d req.JoinCafe) JoinMemberDto {
	return JoinMemberDto{Nickname: d.Nickname}
}

type MemberInfoDto struct {
	Id        int    `json:"member_id,omitempty"`
	UserId    int    `json:"user_id"`
	NickName  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertToTime(s string) (time.Time, error) {
	t, err := time.ParseInLocation(time.RFC3339, s, koreaZone)
	return t, err
}
