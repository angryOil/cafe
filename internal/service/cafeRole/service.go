package cafeRole

import (
	"cafe/internal/cli/cafeRole"
	req2 "cafe/internal/cli/cafeRole/req"
	cafeRole2 "cafe/internal/domain/cafeRole"
	page2 "cafe/internal/page"
	"cafe/internal/service/cafeRole/req"
	"cafe/internal/service/cafeRole/res"
	"context"
)

type Service struct {
	r cafeRole.Requester
}

func NewService(r cafeRole.Requester) Service {
	return Service{r: r}
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.GetList, int, error) {
	domains, total, err := s.r.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.GetList{}, 0, err
	}
	dto := make([]res.GetList, len(domains))
	for i, d := range domains {
		v := d.ToInfo()
		dto[i] = res.GetList{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
		}
	}
	return dto, total, err
}

func (s Service) Create(ctx context.Context, c req.Create) error {
	err := cafeRole2.NewBuilder().
		CafeId(c.CafeId).
		Name(c.Name).
		Description(c.Description).
		Build().ValidCreate()
	if err != nil {
		return err
	}
	err = s.r.Create(ctx, req2.Create{
		CafeId:      c.CafeId,
		Name:        c.Name,
		Description: c.Description,
	})
	return err
}

func (s Service) Patch(ctx context.Context, p req.Patch) error {
	err := cafeRole2.NewBuilder().
		Id(p.Id).
		CafeId(p.CafeId).
		Name(p.Name).
		Description(p.Description).
		Build().ValidUpdate()
	if err != nil {
		return err
	}
	err = s.r.Patch(ctx, req2.Patch{
		Id:          p.Id,
		CafeId:      p.CafeId,
		Name:        p.Name,
		Description: p.Description,
	})
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := s.r.Delete(ctx, cafeId, roleId)
	return err
}
