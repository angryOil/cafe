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

func (s CafeService) Update(ctx context.Context, reqDomain domain.Cafe) error {
	err := reqDomain.ValidCafeFiled()
	if err != nil {
		return err
	}

	err = s.repo.Save(ctx,
		reqDomain.OwnerId, reqDomain.Id,
		func(results []domain.Cafe) (domain.Cafe, error) {
			if len(results) == 0 {
				return domain.Cafe{}, errors.New("this cafe is not exists")
			}
			return results[0], nil
		},
		func(findDomain domain.Cafe) domain.Cafe {
			findDomain.Name = reqDomain.Name
			findDomain.Description = reqDomain.Description
			return findDomain
		},
		func(cafe domain.Cafe) error {
			// error가 나올일이 없다고 생각하지만 혹시 error가 나온다면 크리티컬하기 때문에 한번더 확인
			if cafe.Id == 0 {
				return errors.New("cafe id is zero")
			}
			if cafe.OwnerId == 0 {
				return errors.New("owner id is zero")
			}
			if cafe.Name == "" {
				return errors.New("cafe name is empty")
			}
			return nil
		},
	)
	return err
}

func (s CafeService) GetListByIds(ctx context.Context, ids []int) ([]domain.Cafe, error) {
	cDomains, err := s.repo.GetCafesByCafeIds(ctx, ids)
	return cDomains, err
}

func (s CafeService) CheckIsMine(ctx context.Context, userId int, cafeId int) (bool, error) {
	ok, err := s.repo.IsExistsByUserIdCafeId(ctx, userId, cafeId)
	return ok, err
}

func (s CafeService) IsExistsCafe(ctx context.Context, cafeId int) (bool, error) {
	ok, err := s.repo.IsExistsByCafeId(ctx, cafeId)
	return ok, err
}
