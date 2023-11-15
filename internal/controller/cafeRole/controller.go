package cafeRole

import (
	"cafe/internal/controller/cafeRole/req"
	"cafe/internal/controller/cafeRole/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/cafeRole"
	req2 "cafe/internal/service/cafeRole/req"
	"context"
)

type Controller struct {
	s cafeRole.Service
}

func NewController(s cafeRole.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.RoleDto, int, error) {
	list, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.RoleDto{}, 0, err
	}
	dto := make([]res.RoleDto, len(list))
	for i, l := range list {
		dto[i] = res.RoleDto{
			Id:          l.Id,
			Name:        l.Name,
			Description: l.Description,
		}
	}
	return dto, total, nil
}

func (c Controller) Create(ctx context.Context, cafeId int, roleDto req.CreateRoleDto) error {
	err := c.s.Create(ctx, req2.Create{
		CafeId:      cafeId,
		Name:        roleDto.Name,
		Description: roleDto.Description,
	})
	return err
}

func (c Controller) Patch(ctx context.Context, cafeId int, roleId int, d req.PatchRoleDto) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:          roleId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := c.s.Delete(ctx, cafeId, roleId)
	return err
}
