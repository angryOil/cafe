package service

import (
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"cafe/internal/repository"
	"context"
	"errors"
	"log"
)

type CafeService struct {
	repo repository.CafeRepository
}

func NewService(repo repository.CafeRepository) CafeService {
	return CafeService{repo: repo}
}

func (s CafeService) CreateCafe(ctx context.Context, d domain.Cafe) error {
	err := s.repo.Create(ctx, d)
	return err
}

func (s CafeService) GetCafes(ctx context.Context, reqPage page2.ReqPage) ([]domain.Cafe, int, error) {
	cafes, count, err := s.repo.GetCafes(ctx, reqPage)
	if err != nil {
		log.Println("getCafes err: ", err)
		return cafes, count, errors.New("internal server error")
	}
	return cafes, count, nil
}

func (s CafeService) GetDetail(ctx context.Context, id int) (domain.Cafe, error) {
	if id == 0 {
		return domain.Cafe{}, errors.New("id is zero")
	}

	resuls, err := s.repo.GetDetail(ctx, id)
	if err != nil {
		log.Println("getDetail err: ", err)
		return domain.Cafe{}, errors.New("internal server error")
	}
	if len(resuls) == 0 {
		return domain.Cafe{}, nil
	}
	return resuls[0], nil
}
