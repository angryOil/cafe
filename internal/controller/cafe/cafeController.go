package cafe

import (
	"cafe/internal/controller/cafe/req"
	"cafe/internal/controller/cafe/res"
	page2 "cafe/internal/page"
	"cafe/internal/service/cafe"
	"cafe/internal/service/cafe/request"
	"cafe/internal/service/cafe/response"
	"context"
	"errors"
)

type Controller struct {
	s cafe.Service
}

func NewCafeController(s cafe.Service) Controller {
	return Controller{s: s}
}

func (c Controller) CreateCafe(ctx context.Context, dto req.CreateCafeDto) error {
	// userId 를 검사한다
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return errors.New("user id is not valid")
	}
	// 카페를 저장한다
	err := c.s.CreateCafe(ctx, request.CreateCafe{
		OwnerId:     userId,
		Name:        dto.Name,
		Description: dto.Description,
	})
	// 에러 처리를 한다
	return err
}

func (c Controller) GetCafes(ctx context.Context, reqPage page2.ReqPage) ([]res.CafeListDto, int, error) {
	// 카페 리스트, 리스트 수를 반환한다
	cafes, count, err := c.s.GetCafes(ctx, reqPage)
	if err != nil {
		return nil, 0, err
	}

	return c.convertToCafeListDto(cafes), count, err
}

func (c Controller) convertToCafeListDto(cafes []response.GetCafes) []res.CafeListDto {
	dtoList := make([]res.CafeListDto, len(cafes))

	for i, cf := range cafes {
		dto := res.CafeListDto{
			Id:   cf.Id,
			Name: cf.Name,
		}
		dtoList[i] = dto
	}
	return dtoList
}

func (c Controller) GetDetail(ctx context.Context, id int) (res.CafeDetailDto, error) {
	result, err := c.s.GetDetail(ctx, id)
	if err != nil {
		return res.CafeDetailDto{}, err
	}
	return res.ToDetailDto(result.ToDetail()), err
}

const (
	InvalidUserId = "invalid user id"
)

func (c Controller) Update(ctx context.Context, dto req.UpdateCafeDto, cafeId int) error {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return errors.New(InvalidUserId)
	}
	err := c.s.Update(ctx, request.Update{
		Id:          cafeId,
		OwnerId:     userId,
		Name:        dto.Name,
		Description: dto.Description,
	})
	return err
}

func (c Controller) GetCafesByCafeIds(ctx context.Context, ids []int) ([]res.CafeListDto, error) {
	if len(ids) == 0 {
		return []res.CafeListDto{}, nil
	}
	r, err := c.s.GetListByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	return c.convertCafeListDtoList(r), nil
}

func (c Controller) convertCafeListDtoList(r []response.GetListByIds) []res.CafeListDto {
	result := make([]res.CafeListDto, len(r))
	for i, byIds := range r {
		result[i] = res.CafeListDto{
			Id:   byIds.Id,
			Name: byIds.Name,
		}
	}
	return result
}

func (c Controller) CheckIsMine(ctx context.Context, userId int, cafeId int) (bool, error) {
	isMine, err := c.s.CheckIsMine(ctx, userId, cafeId)
	return isMine, err
}

func (c Controller) GetOwnerId(ctx context.Context, cafeId int) (int, error) {
	ownerDto, err := c.s.GetOwnerId(ctx, cafeId)
	if err != nil {
		return 0, err
	}
	return ownerDto.Id, nil
}

func (c Controller) IsExistsCafe(ctx context.Context, cafeId int) (bool, error) {
	ok, err := c.s.IsExistsCafe(ctx, cafeId)
	return ok, err
}
