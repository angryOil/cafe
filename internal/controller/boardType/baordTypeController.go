package boardType

import (
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
