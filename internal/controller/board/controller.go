package board

import (
	req2 "cafe/internal/controller/board/req"
	"cafe/internal/controller/board/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/board"
	"cafe/internal/service/board/req"
	"context"
)

type Controller struct {
	s board.Service
}

func (c Controller) GetList(ctx context.Context, cafeId int, boardType int, writer int, reqPage page2.ReqPage) ([]res.Info, int, error) {
	list, total, err := c.s.GetList(ctx, req.GetList{
		CafeId:    cafeId,
		BoardType: boardType,
		Writer:    writer,
	}, reqPage)
	if err != nil {
		return []res.Info{}, 0, err
	}
	dto := make([]res.Info, len(list))
	for i, l := range list {
		dto[i] = res.Info{
			Id:            l.Id,
			BoardType:     l.BoardType,
			Writer:        l.Writer,
			Title:         l.Title,
			CreatedAt:     l.CreatedAt,
			LastUpdatedAt: l.LastUpdatedAt,
		}
	}
	return dto, total, nil
}

func (c Controller) Create(ctx context.Context, cafeId int, boardType int, memberId int, create req2.Create) error {
	err := c.s.Create(ctx, req.Create{
		Writer:    memberId,
		CafeId:    cafeId,
		BoardType: boardType,
		Title:     create.Title,
		Content:   create.Content,
	})
	return err
}

func NewController(s board.Service) Controller {
	return Controller{s: s}
}
