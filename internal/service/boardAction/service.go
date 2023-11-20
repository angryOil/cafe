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
	}, nil
}

func (s Service) Create(ctx context.Context, c req2.Create) error {
	cafeId, boardType := c.CafeId, c.BoardType
	readRoles, createRoles := c.ReadRoles, c.CreateRoles

	err := boardAction2.NewBuilder().
		CafeId(cafeId).
		BoardTypeId(boardType).
		ReadRoles(readRoles).
		CreateRoles(createRoles).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.r.Create(ctx, req.Create{
		CafeId:      cafeId,
		BoardType:   boardType,
		ReadRoles:   readRoles,
		CreateRoles: createRoles,
	})
	return err
}

func (s Service) Patch(ctx context.Context, p req2.Patch) error {
	id, cafeId, boardTypeId := p.Id, p.CafeId, p.BoardTypeId
	readRoles, createRoles := p.ReadRoles, p.CreateRoles

	err := boardAction2.NewBuilder().
		Id(id).
		CafeId(cafeId).
		BoardTypeId(boardTypeId).
		ReadRoles(readRoles).
		CreateRoles(createRoles).
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
	})
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, boardTypeId int, id int) error {
	err := s.r.Delete(ctx, cafeId, boardTypeId, id)
	return err
}
