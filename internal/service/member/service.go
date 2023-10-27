package member

import (
	"cafe/internal/cli/member"
	"cafe/internal/domain"
	"cafe/internal/page"
	"context"
	"errors"
)

type Service struct {
	r member.Requester
}

func (s Service) GetMemberInfo(ctx context.Context, cafeId int, userId int) (domain.Member, error) {
	if userId == 0 {
		return domain.Member{}, errors.New("invalid user id")
	}
	if cafeId == 0 {
		return domain.Member{}, errors.New("invalid cafe id")
	}

	md, err := s.r.GetCafeMyInfo(ctx, cafeId, userId)
	return md, err
}

func (s Service) GetCafeIdsAndTotal(ctx context.Context, userId int, reqPage page.ReqPage) (domain.IdsTotalDomain, error) {
	id, err := s.r.GetCafeIdsAndTotalByUserId(ctx, userId, reqPage)
	return id, err
}

func (s Service) JoinCafe(ctx context.Context, d domain.Member) error {
	err := s.r.JoinCafe(ctx, d)
	return err
}

func NewService(r member.Requester) Service {
	return Service{r: r}
}
