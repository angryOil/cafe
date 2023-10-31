package cafeRole

import (
	"cafe/internal/controller/cafeRole/req"
	"cafe/internal/controller/cafeRole/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/role"
	"context"
)

type Controller struct {
	s role.Service
}

func NewController(s role.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.RoleDto, int, error) {
	domains, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.RoleDto{}, 0, err
	}
	return res.ToDtoList(domains), total, nil
}

func (c Controller) Create(ctx context.Context, cafeId int, roleDto req.CreateRoleDto) error {
	d := roleDto.ToDomain(cafeId)
	err := c.s.Create(ctx, d)
	return err
}

func (c Controller) Patch(ctx context.Context, cafeId int, roleId int, roleDto req.PatchRoleDto) error {
	d := roleDto.ToDomain(cafeId, roleId)
	err := c.s.Patch(ctx, d)
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := c.s.Delete(ctx, cafeId, roleId)
	return err
}
