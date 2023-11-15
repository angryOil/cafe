package boardAction

import (
	"cafe/internal/controller/boardAction/res"
	"cafe/internal/service/boardAction"
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
