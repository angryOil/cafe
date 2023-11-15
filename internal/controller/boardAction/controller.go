package boardAction

import (
	"cafe/internal/controller/boardAction/req"
	"cafe/internal/controller/boardAction/res"
	"cafe/internal/service/boardAction"
	req2 "cafe/internal/service/boardAction/req"
	"context"
)

type Controller struct {
	s boardAction.Service
}

func NewController(s boardAction.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetInfo(ctx context.Context, cafeId, boardTypeId int) (res.GetInfo, error) {
	d, err := c.s.GetInfo(ctx, cafeId, boardTypeId)
	if err != nil {
		return res.GetInfo{}, err
	}
	return res.GetInfo{
		Id:          d.Id,
		CafeId:      d.CafeId,
		BoardTypeId: d.BoardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
	}, nil
}

func (c Controller) Create(ctx context.Context, cafeId int, boardTypeId int, d req.Create) error {
	err := c.s.Create(ctx, req2.Create{
		CafeId:      cafeId,
		BoardType:   boardTypeId,
		ReadRoles:   d.ReadRoles,
		CreateRoles: d.CreateRoles,
		UpdateRoles: d.UpdateRoles,
		UpdateAble:  d.UpdateAble,
		DeleteRoles: d.DeleteRoles,
	})
	return err
}

func (c Controller) Patch(ctx context.Context, cafeId int, boardTypeId int, id int, p req.Patch) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:          id,
		CafeId:      cafeId,
		BoardTypeId: boardTypeId,
		ReadRoles:   p.ReadRoles,
		CreateRoles: p.CreateRoles,
		UpdateRoles: p.UpdateRoles,
		UpdateAble:  p.UpdateAble,
		DeleteRoles: p.DeleteRoles,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, boardTypeId int, id int) error {
	err := c.s.Delete(ctx, cafeId, boardTypeId, id)
	return err
}
