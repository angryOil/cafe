package member

import (
	"cafe/internal/controller/member/res"
	"cafe/internal/service/member"
	"context"
)

type Controller struct {
	s member.Service
}

func (c Controller) GetMemberInfo(ctx context.Context, cafeId int, userId int) (res.MemberInfoDto, error) {
	md, err := c.s.GetMemberInfo(ctx, cafeId, userId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	return res.ToMemberInfoDto(md), err
}

func NewController(s member.Service) Controller {
	return Controller{s: s}
}
