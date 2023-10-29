package ban

import (
	"cafe/internal/domain"
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
