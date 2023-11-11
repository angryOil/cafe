package memberRole

import (
	req2 "cafe/internal/controller/memberRole/req"
	"cafe/internal/controller/memberRole/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/memberRole"
	"cafe/internal/service/memberRole/req"
	"context"
)

type Controller struct {
	s memberRole.Service
}

func NewController(s memberRole.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.DetailDto, int, error) {
	list, total, err := c.s.GetRolesByCafeId(ctx, cafeId, reqPage)
	if err != nil {
		return []res.DetailDto{}, 0, err
	}
	dto := make([]res.DetailDto, len(list))
	for i, l := range list {
		dto[i] = res.DetailDto{
			Id:          l.Id,
			CafeRoleIds: l.CafeRoleIds,
			MemberId:    l.MemberId,
		}
	}
	return dto, total, err
}

func (c Controller) GetOneMemberRoles(ctx context.Context, cafeId int, memberId int) (res.MemberRoleDto, error) {
	d, err := c.s.GetOneMemberRoles(ctx, cafeId, memberId)
	if err != nil {
		return res.MemberRoleDto{}, err
	}
	return res.MemberRoleDto{
		Id:          d.Id,
		CafeRoleIds: d.CafeRoleIds,
	}, nil
}

func (c Controller) PutRole(ctx context.Context, id, cafeId, memberId int, putDto req2.PutMemberRoleDto) error {
	err := c.s.PutRole(ctx, req.PutRole{
		Id:          id,
		CafeId:      cafeId,
		MemberId:    memberId,
		CafeRoleIds: putDto.CafeRoleIds,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, memberId int, mRoleId int) error {
	err := c.s.Delete(ctx, cafeId, memberId, mRoleId)
	return err
}

func (c Controller) Create(ctx context.Context, cafeId int, memberId int, dto req2.CreateRoleDto) error {
	err := c.s.CreateRole(ctx, req.CreateRole{
		CafeId:      cafeId,
		MemberId:    memberId,
		CafeRoleIds: dto.CafeRoleIds,
	})
	return err
}
