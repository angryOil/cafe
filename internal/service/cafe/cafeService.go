package cafe

import (
	"cafe/internal/domain/cafe"
	"cafe/internal/domain/cafe/vo"
	page2 "cafe/internal/page"
	"cafe/internal/repository"
	request2 "cafe/internal/repository/request"
	"cafe/internal/service/cafe/request"
	"cafe/internal/service/cafe/response"
	"context"
	"errors"
	"log"
	"time"
)

type Service struct {
	repo repository.CafeRepository
}

func NewService(repo repository.CafeRepository) Service {
	return Service{repo: repo}
}

func (s Service) CreateCafe(ctx context.Context, req request.CreateCafe) error {
	// 생성일시 할당
	createdAt := time.Now()

	// cafe 생성 검증
	c := cafe.NewCafeBuilder().
		Name(req.Name).
		Description(req.Description).
		OwnerId(req.OwnerId).
		CreatedAt(createdAt).
		Build()
	err := c.ValidCreate()
	if err != nil {
		return err
	}

	// cafe 생성
	err = s.repo.Create(ctx, request2.CreateCafe{
		OwnerId:     req.OwnerId,
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   createdAt,
	})
	return err
}

func (s Service) GetCafes(ctx context.Context, reqPage page2.ReqPage) ([]response.GetCafes, int, error) {
	cafes, count, err := s.repo.GetCafes(ctx, reqPage)
	if err != nil {
		log.Println("getCafes err: ", err)
		return nil, 0, errors.New("internal server error")
	}

	dto := make([]response.GetCafes, len(cafes))
	for i, cafe := range cafes {
		vo := cafe.ToCafeListInfo()
		dto[i] = response.GetCafes{
			Id:   vo.Id,
			Name: vo.Name,
		}
	}
	return dto, count, nil
}

func (s Service) GetDetail(ctx context.Context, id int) (cafe.Cafe, error) {
	if id == 0 {
		return cafe.NewCafeBuilder().Build(), errors.New("id is zero")
	}

	results, err := s.repo.GetDetail(ctx, id)
	if err != nil {
		log.Println("getDetail err: ", err)
		return cafe.NewCafeBuilder().Build(), errors.New("internal server error")
	}
	if len(results) == 0 {
		return cafe.NewCafeBuilder().Build(), nil
	}
	return results[0], nil
}

const (
	InvalidCafeId       = "invalid cafe id"
	InternalServerError = "internal server error"
	NotFoundOwnerId     = "not found ownerId"
)

func (s Service) GetOwnerId(ctx context.Context, id int) (response.OwnerId, error) {
	if id < 1 {
		return response.OwnerId{}, errors.New(InvalidCafeId)
	}

	cafes, err := s.repo.GetOwnerIds(ctx, id)
	if err != nil {
		log.Println("getDetail err: ", err)
		return response.OwnerId{}, errors.New(InternalServerError)
	}
	if len(cafes) != 1 {
		return response.OwnerId{}, errors.New(NotFoundOwnerId)
	}
	return response.OwnerId{Id: cafes[0].GetOwnerId()}, nil
}

func (s Service) Update(ctx context.Context, req request.Update) error {
	var id, ownerId = req.Id, req.OwnerId
	var name, description = req.Name, req.Description

	err := cafe.NewCafeBuilder().
		Id(id).
		OwnerId(ownerId).
		Name(name).
		Description(description).
		Build().
		ValidCafeFiled()
	if err != nil {
		return err
	}

	err = s.repo.Save(ctx,
		ownerId, id,
		func(results []cafe.Cafe) (cafe.Cafe, error) {
			if len(results) == 0 {
				return cafe.NewCafeBuilder().Build(), errors.New("this cafe is not exists")
			}
			return results[0], nil
		},
		func(findDomain cafe.Cafe) (vo.UpdateCafe, error) {
			updatedCafe := findDomain.Update(name, description)
			err := updatedCafe.VerifyUpdate()
			if err != nil {
				return vo.UpdateCafe{}, err
			}
			return findDomain.UpdateCafeInfo(), nil
		},
		func(cafe cafe.Cafe) error {
			return cafe.VerifyUpdate()
		},
	)
	return err
}

func (s Service) GetListByIds(ctx context.Context, ids []int) ([]response.GetListByIds, error) {
	cDomains, err := s.repo.GetCafesByCafeIds(ctx, ids)
	dto := make([]response.GetListByIds, len(cDomains))
	for i, cDomain := range cDomains {
		vo := cDomain.ToCafeListInfo()
		dto[i] = response.GetListByIds{
			Id:   vo.Id,
			Name: vo.Name,
		}
	}
	return dto, err
}

func (s Service) CheckIsMine(ctx context.Context, userId int, cafeId int) (bool, error) {
	ok, err := s.repo.IsExistsByUserIdCafeId(ctx, userId, cafeId)
	return ok, err
}

func (s Service) IsExistsCafe(ctx context.Context, cafeId int) (bool, error) {
	ok, err := s.repo.IsExistsByCafeId(ctx, cafeId)
	return ok, err
}
