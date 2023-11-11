package ban

import (
	"cafe/internal/domain/ban"
	page2 "cafe/internal/page"
	ban2 "cafe/internal/repository/ban"
	req2 "cafe/internal/repository/ban/req"
	"cafe/internal/service/ban/req"
	"cafe/internal/service/ban/res"
	"context"
	"time"
)

type Service struct {
	repo ban2.BanRepository
}

func NewService(repo ban2.BanRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateBan(ctx context.Context, c req.CreateBan) error {
	userId, memberId, cafeId := c.UserId, c.MemberId, c.CafeId
	description := c.Description
	createdAt := time.Now()

	err := ban.NewBuilder().
		UserId(userId).
		MemberId(memberId).
		CafeId(cafeId).
		Description(description).
		CreatedAt(createdAt).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, req2.Create{
		UserId:      userId,
		MemberId:    memberId,
		CafeId:      cafeId,
		Description: description,
		CreatedAt:   createdAt,
	})
	return err
}

func (s Service) GetBanListAndCountByUserId(ctx context.Context, userId int, reqPage page2.ReqPage) ([]res.GetBanListAndCountByUserId, int, error) {
	domains, count, err := s.repo.GetListCountByUserId(ctx, userId, reqPage)
	result := make([]res.GetBanListAndCountByUserId, len(domains))
	for i, d := range domains {
		v := d.ToInfo()
		result[i] = res.GetBanListAndCountByUserId{
			Id:          v.Id,
			CafeId:      v.CafeId,
			Description: v.Description,
		}
	}
	return result, count, err
}

func (s Service) GetBanListByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.GetBanListByCafeId, int, error) {
	domains, count, err := s.repo.GetListCountByCafeId(ctx, cafeId, reqPage)
	result := make([]res.GetBanListByCafeId, len(domains))
	for i, d := range domains {
		v := d.ToInfo()
		result[i] = res.GetBanListByCafeId{
			Id:          v.Id,
			MemberId:    v.MemberId,
			Description: v.Description,
		}
	}
	return result, count, err
}
