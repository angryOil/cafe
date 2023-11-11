package boardType

import (
	"cafe/internal/controller/boardType/req"
	"cafe/internal/controller/boardType/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/boardType"
	req2 "cafe/internal/service/boardType/req"
	"context"
)

type Controller struct {
	s boardType.Service
}

func NewController(s boardType.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.BoardTypeDto, int, error) {
	list, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.BoardTypeDto{}, 0, err
	}
	dto := make([]res.BoardTypeDto, len(list))
	for i, l := range list {
		dto[i] = res.BoardTypeDto{
			Id:          l.Id,
			Name:        l.Name,
			Description: l.Description,
		}
	}
	return dto, total, nil
}

func (c Controller) Create(ctx context.Context, cafeId, ownerId int, d req.CreateBoardTypeDto) error {
	err := c.s.Create(ctx, req2.Create{
		Name:        d.Name,
		Description: d.Description,
		CafeId:      cafeId,
		OwnerId:     ownerId,
	})
	return err
}

func (c Controller) Patch(ctx context.Context, cafeId, typeId int, d req.PatchBoardTypeDto) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:          typeId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, typeId int) error {
	err := c.s.Delete(ctx, cafeId, typeId)
	return err
}
