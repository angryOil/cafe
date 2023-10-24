package controller

import (
	"cafe/internal/controller/req"
	"cafe/internal/controller/res"
	page2 "cafe/internal/page"
	"cafe/internal/service"
	"context"
	"errors"
)

type CafeController struct {
	s service.CafeService
}

func NewCafeController(s service.CafeService) CafeController {
	return CafeController{s: s}
}

func (c CafeController) CreateCafe(ctx context.Context, dto req.CreateCafeDto) error {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return errors.New("user id is not valid")
	}
	cafeDomain, err := dto.ToDomain(userId)
	if err != nil {
		return err
	}
	err = c.s.CreateCafe(ctx, cafeDomain)
	return err
}

func (c CafeController) GetCafes(ctx context.Context, reqPage page2.ReqPage) ([]res.CafeListDto, int, error) {
	cafes, count, err := c.s.GetCafes(ctx, reqPage)
	dtos := res.ToListDtoList(cafes)
	return dtos, count, err
}

func (c CafeController) GetDetail(ctx context.Context, id int) (res.CafeDetailDto, error) {
	result, err := c.s.GetDetail(ctx, id)
	if err != nil {
		return res.CafeDetailDto{}, err
	}
	return res.ToDetailDto(result), err
}
