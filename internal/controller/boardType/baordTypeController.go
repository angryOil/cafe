package boardType

import (
	"cafe/internal/controller/boardType/req"
	"cafe/internal/controller/boardType/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/boardType"
	"context"
)

type Controller struct {
	s boardType.Service
}

func NewController(s boardType.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.BoardTypeDto, int, error) {
	domains, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.BoardTypeDto{}, 0, err
	}
	return res.ToBoardTypeDtoList(domains), total, nil
}

func (c Controller) Create(ctx context.Context, cafeId, ownerId int, d req.CreateBoardTypeDto) error {
	typeDomain := d.ToDomain(cafeId, ownerId)
	err := c.s.Create(ctx, typeDomain)
	return err
}

func (c Controller) Patch(ctx context.Context, cafeId, typeId int, d req.PatchBoardTypeDto) error {
	tyDomain := d.ToDomain(cafeId, typeId)
	err := c.s.Patch(ctx, tyDomain)
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, typeId int) error {
	err := c.s.Delete(ctx, cafeId, typeId)
	return err
}
