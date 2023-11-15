package boardAction

import (
	"cafe/internal/cli/boardAction"
	"cafe/internal/cli/boardAction/req"
	boardAction2 "cafe/internal/domain/boardAction"
	req2 "cafe/internal/service/boardAction/req"
	"cafe/internal/service/boardAction/res"
	"context"
)

type Service struct {
	r boardAction.Requester
}

func NewService(r boardAction.Requester) Service {
	return Service{r: r}
}

func (s Service) GetInfo(ctx context.Context, cafeId, boardTypeId int) (res.GetInfo, error) {
	d, err := s.r.GetInfo(ctx, cafeId, boardTypeId)
	if err != nil {
		return res.GetInfo{}, err
	}
	v := d.ToInfo()
	return res.GetInfo{
		Id:          v.Id,
		CafeId:      v.CafeId,
		BoardTypeId: v.BoardTypeId,
		ReadRoles:   v.ReadRoles,
		CreateRoles: v.CreateRoles,
		UpdateRoles: v.UpdateRoles,
		UpdateAble:  v.UpdateAble,
		DeleteRoles: v.DeleteRoles,
	}, nil
}

func (s Service) Create(ctx context.Context, c req2.Create) error {
	cafeId, boardType := c.CafeId, c.BoardType
	readRoles, createRoles, updateRoles, deleteRoles := c.ReadRoles, c.CreateRoles, c.UpdateRoles, c.DeleteRoles
	updateAble := c.UpdateAble

	err := boardAction2.NewBuilder().
		CafeId(cafeId).
		BoardTypeId(boardType).
		ReadRoles(readRoles).
		CreateRoles(createRoles).
		UpdateRoles(updateRoles).
		UpdateAble(updateAble).
		DeleteRoles(deleteRoles).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.r.Create(ctx, req.Create{
		CafeId:      cafeId,
		BoardType:   boardType,
		ReadRoles:   readRoles,
		CreateRoles: createRoles,
		UpdateRoles: updateRoles,
		UpdateAble:  updateAble,
		DeleteRoles: deleteRoles,
	})
	return err
}

func (s Service) Patch(ctx context.Context, p req2.Patch) error {
	id, cafeId, boardTypeId := p.Id, p.CafeId, p.BoardTypeId
	readRoles, createRoles, updateRoles, deleteRoles := p.ReadRoles, p.CreateRoles, p.UpdateRoles, p.DeleteRoles
	updateAble := p.UpdateAble

	err := boardAction2.NewBuilder().
		Id(id).
		CafeId(cafeId).
		BoardTypeId(boardTypeId).
		ReadRoles(readRoles).
		CreateRoles(createRoles).
		UpdateRoles(updateRoles).
		UpdateAble(updateAble).
		DeleteRoles(deleteRoles).
		Build().ValidUpdate()
	if err != nil {
		return err
	}

	err = s.r.Patch(ctx, req.Patch{
		Id:          id,
		CafeId:      cafeId,
		BoardTypeId: boardTypeId,
		ReadRoles:   readRoles,
		CreateRoles: createRoles,
		UpdateRoles: updateRoles,
		UpdateAble:  updateAble,
		DeleteRoles: deleteRoles,
	})
	return err
}
