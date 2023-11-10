package memberRole

import (
	"cafe/internal/cli/memberRole"
	req2 "cafe/internal/cli/memberRole/req"
	memberRole2 "cafe/internal/domain/memberRole"
	page2 "cafe/internal/page"
	"cafe/internal/service/memberRole/req"
	"cafe/internal/service/memberRole/res"
	"context"
)

type Service struct {
	r memberRole.Requester
}

func NewService(r memberRole.Requester) Service {
	return Service{r: r}
}

func (s Service) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]res.GetRolesByCafeId, int, error) {
	domains, total, err := s.r.GetRolesByCafeId(ctx, cafeId, reqPage)
	if err != nil {
		return []res.GetRolesByCafeId{}, 0, err
	}
	dto := make([]res.GetRolesByCafeId, len(domains))
	for i, d := range domains {
		v := d.ToDetail()
		dto[i] = res.GetRolesByCafeId{
			Id:          v.Id,
			CafeRoleIds: v.CafeRoleIds,
			MemberId:    v.MemberId,
		}
	}
	return dto, total, err
}

func (s Service) GetOneMemberRoles(ctx context.Context, cafeId int, memberId int) (res.GetOneMemberRoles, error) {
	d, err := s.r.GetOneMemberRoles(ctx, cafeId, memberId)
	v := d.ToInfo()
	return res.GetOneMemberRoles{
		Id:          v.Id,
		CafeRoleIds: v.CafeRoleIds,
	}, err
}

func (s Service) PutRole(ctx context.Context, p req.PutRole) error {
	err := s.r.PutRole(ctx, req2.PutRole{
		Id:          p.Id,
		CafeId:      p.CafeId,
		MemberId:    p.MemberId,
		CafeRoleIds: p.CafeRoleIds,
	})
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, memberId int, mRoleId int) error {
	err := s.r.Delete(ctx, cafeId, memberId, mRoleId)
	return err
}

func (s Service) CreateRole(ctx context.Context, c req.CreateRole) error {
	err := memberRole2.NewBuilder().
		CafeId(c.CafeId).
		MemberId(c.MemberId).
		CafeRoleIds(c.CafeRoleIds).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.r.Create(ctx, req2.CreateRole{
		CafeId:      c.CafeId,
		MemberId:    c.MemberId,
		CafeRoleIds: c.CafeRoleIds,
	})
	return err
}
