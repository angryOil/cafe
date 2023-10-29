package ban

import (
	"cafe/internal/controller/ban/req"
	"cafe/internal/controller/ban/res"
	page2 "cafe/internal/page"
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

func (c Controller) GetMyBanListAndCount(ctx context.Context, userId int, reqPage page2.ReqPage) ([]res.BanListDto, int, error) {
	domains, count, err := c.s.GetBanListAndCountByUserId(ctx, userId, reqPage)
	if err != nil {
		return []res.BanListDto{}, 0, err
	}
	return res.ToBandListDtoList(domains), count, nil
}

func NewController(s ban.Service) Controller {
	return Controller{s: s}
}
