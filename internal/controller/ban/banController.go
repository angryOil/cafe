package ban

import (
	"cafe/internal/controller/ban/req"
	"cafe/internal/service/ban"
	"context"
)

type Controller struct {
	s ban.Service
}

func (c Controller) CreateBan(ctx context.Context, userId int, cafeId int, dto req.CreateBanDto) error {
	bDomain := dto.ToDomain(userId, cafeId)
	err := c.s.CreateBan(ctx, bDomain)
	return err
}

func NewController(s ban.Service) Controller {
	return Controller{s: s}
}
