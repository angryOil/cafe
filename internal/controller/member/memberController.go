package member

import (
	"cafe/internal/controller/member/req"
	"cafe/internal/controller/member/res"
	"cafe/internal/page"
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

func (c Controller) GetMyCafeIds(ctx context.Context, userId int, reqPage page.ReqPage) (res.IdsTotalDto, error) {
	iTDto, err := c.s.GetCafeIdsAndTotal(ctx, userId, reqPage)
	if err != nil {
		return res.IdsTotalDto{}, err
	}
	iTDomain := res.ToIdsTotalDto(iTDto)
	return iTDomain, nil
}

func (c Controller) JoinCafe(ctx context.Context, userId, cafeId int, dto req.JoinMemberDto) error {
	d := dto.ToDomain(userId, cafeId)
	err := c.s.JoinCafe(ctx, d)
	return err
}
