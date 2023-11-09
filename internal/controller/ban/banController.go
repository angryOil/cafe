package ban

import (
	"cafe/internal/controller/ban/req"
	"cafe/internal/controller/ban/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/ban"
	req2 "cafe/internal/service/ban/req"
	"context"
)

type Controller struct {
	s ban.Service
}

func NewController(s ban.Service) Controller {
	return Controller{s: s}
}

func (c Controller) CreateBan(ctx context.Context, userId int, cafeId int, dto req.CreateBanDto) error {
	err := c.s.CreateBan(ctx, req2.CreateBan{
		UserId:      userId,
		MemberId:    dto.MemberId,
		CafeId:      cafeId,
		Description: dto.Description,
	})
	return err
}

func (c Controller) GetMyBanListAndCount(ctx context.Context, userId int, reqPage page2.ReqPage) ([]res.BanListDto, int, error) {
	list, count, err := c.s.GetBanListAndCountByUserId(ctx, userId, reqPage)
	if err != nil {
		return []res.BanListDto{}, 0, err
	}
	result := make([]res.BanListDto, len(list))
	for i, l := range list {
		result[i] = res.BanListDto{
			Id:          l.Id,
			CafeId:      l.CafeId,
			Description: l.Description,
		}
	}
	return result, count, nil
}

func (c Controller) GetBanListByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.BanAdminListDto, int, error) {
	list, count, err := c.s.GetBanListByCafeId(ctx, cafeId, reqPage)
	if err != nil {
		return []res.BanAdminListDto{}, 0, err
	}
	result := make([]res.BanAdminListDto, len(list))
	for i, l := range list {
		result[i] = res.BanAdminListDto{
			Id:          l.Id,
			MemberId:    l.MemberId,
			Description: l.Description,
		}
	}
	return result, count, nil
}
