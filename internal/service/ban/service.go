package ban

import (
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"cafe/internal/repository"
	"context"
	"errors"
)

type Service struct {
	repo repository.BanRepository
}

func NewService(repo repository.BanRepository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateBan(ctx context.Context, bDomain domain.Ban) error {
	if bDomain.MemberId == 0 {
		return errors.New("invalid member id")
	}
	err := s.repo.Create(ctx, bDomain)
	return err
}

func (s Service) GetBanListAndCountByUserId(ctx context.Context, userId int, reqPage page2.ReqPage) ([]domain.Ban, int, error) {
	domains, count, err := s.repo.GetListCountByUserId(ctx, userId, reqPage)
	return domains, count, err
}
