package role

import (
	"cafe/internal/cli/cafeRole"
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"context"
)

type Service struct {
	r cafeRole.Requester
}

func NewService(r cafeRole.Requester) Service {
	return Service{r: r}
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.CafeRole, int, error) {
	domains, total, err := s.r.GetList(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s Service) Create(ctx context.Context, d domain.CafeRole) error {
	err := s.r.Create(ctx, d)
	return err
}

func (s Service) Patch(ctx context.Context, d domain.CafeRole) error {
	err := s.r.Patch(ctx, d)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := s.r.Delete(ctx, cafeId, roleId)
	return err
}
