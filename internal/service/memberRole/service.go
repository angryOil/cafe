package memberRole

import (
	"cafe/internal/cli/memberRole"
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"context"
)

type Service struct {
	r memberRole.Requester
}

func NewService(r memberRole.Requester) Service {
	return Service{r: r}
}

func (s Service) GetRolesByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.MemberRole, int, error) {
	domains, total, err := s.r.GetRolesByCafeId(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s Service) GetOneMemberRoles(ctx context.Context, cafeId int, memberId int) (domain.MemberRole, error) {
	d, err := s.r.GetOneMemberRoles(ctx, cafeId, memberId)
	return d, err
}

func (s Service) PutRole(ctx context.Context, cafeId int, memberId int, d domain.MemberRole) error {
	err := s.r.PutRole(ctx, cafeId, memberId, d)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, memberId int, mRoleId int) error {
	err := s.r.Delete(ctx, cafeId, memberId, mRoleId)
	return err
}
