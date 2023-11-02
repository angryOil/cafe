package memberRole

import (
	req2 "cafe/internal/controller/memberRole/req"
	"cafe/internal/controller/memberRole/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/memberRole"
	"context"
)

type Controller struct {
	s memberRole.Service
}

func NewController(s memberRole.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.DetailDto, int, error) {
	domains, total, err := c.s.GetRolesByCafeId(ctx, cafeId, reqPage)
	if err != nil {
		return []res.DetailDto{}, 0, err
	}
	return res.ToDetailList(domains), total, err
}

func (c Controller) GetOneMemberRoles(ctx context.Context, cafeId int, memberId int) (res.MemberRoleDto, error) {
	d, err := c.s.GetOneMemberRoles(ctx, cafeId, memberId)
	if err != nil {
		return res.MemberRoleDto{}, err
	}
	return res.ToMemberRoleDto(d), nil
}

func (c Controller) PutRole(ctx context.Context, cafeId, memberId int, putDto req2.PutMemberRoleDto) error {
	d := putDto.ToDomain()
	err := c.s.PutRole(ctx, cafeId, memberId, d)
	return err
}
