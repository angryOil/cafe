package boardType

import (
	"cafe/internal/cli/boardType"
	req2 "cafe/internal/cli/boardType/req"
	boardType2 "cafe/internal/domain/boardType"
	page2 "cafe/internal/page"
	"cafe/internal/service/boardType/req"
	"cafe/internal/service/boardType/res"
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
	err := boardType2.NewBuilder().
		CreateBy(c.OwnerId).
		Name(c.Name).
		Description(c.Description).
		CafeId(c.CafeId).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.r.Create(ctx, req2.Create{
		Name:        c.Name,
		Description: c.Description,
		CafeId:      c.CafeId,
		OwnerId:     c.OwnerId,
	})
	return err
}

func (s Service) Patch(ctx context.Context, p req.Patch) error {
	err := boardType2.NewBuilder().
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

func (s Service) Delete(ctx context.Context, cafeId int, typeId int) error {
	err := s.r.Delete(ctx, cafeId, typeId)
	return err
}

func (s Service) GetDetail(ctx context.Context, cafeId int, id int) (res.GetDetail, error) {
	d, err := s.r.GetDetail(ctx, cafeId, id)
	if err != nil {
		return res.GetDetail{}, err
	}
	v := d.ToInfo()
	return res.GetDetail{
		Id:          v.Id,
		Name:        v.Name,
		Description: v.Description,
	}, nil
}
