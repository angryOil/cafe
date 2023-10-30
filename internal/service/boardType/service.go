package boardType

import (
	"cafe/internal/cli/boardType"
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"context"
)

type Service struct {
	r boardType.Requester
}

func NewService(r boardType.Requester) Service {
	return Service{
		r: r,
	}
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.BoardType, int, error) {
	domains, total, err := s.r.GetList(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s Service) Create(ctx context.Context, typeDomain domain.BoardType) error {
	err := s.r.Create(ctx, typeDomain)
	return err
}
