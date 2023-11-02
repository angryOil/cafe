package memberRole

import (
	"cafe/internal/controller/memberRole/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/memberRole"
	"context"
)

type Controller struct {
	s memberRole.Service
}

func NewController(s memberRole.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.DetailDto, int, error) {
	domains, total, err := c.s.GetRolesByCafeId(ctx, cafeId, reqPage)
	if err != nil {
		return []res.DetailDto{}, 0, err
	}
	return res.ToDetailList(domains), total, err
}
