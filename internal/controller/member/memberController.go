package member

import (
	"cafe/internal/controller/member/res"
	"cafe/internal/service/member"
	"context"
)

type Controller struct {
	s member.Service
}

func NewController(s member.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetMemberInfo(ctx context.Context, cafeId int, userId int) (res.MemberInfoDto, error) {
	md, err := c.s.GetMemberInfo(ctx, cafeId, userId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	return res.ToMemberInfoDto(md), err
}

func (c Controller) GetMyCafeIds(ctx context.Context, userId int) (res.IdsTotalDto, error) {
	iTDto, err := c.s.GetCafeIdsAndTotal(ctx, userId)
	if err != nil {
		return res.IdsTotalDto{}, err
	}
	iTDomain := res.ToIdsTotalDto(iTDto)
	return iTDomain, nil
}
