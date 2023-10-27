package member

import (
	"cafe/internal/cli/member"
	"cafe/internal/domain"
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

func (s Service) GetCafeIdsAndTotal(ctx context.Context, userId int) (domain.IdsTotalDomain, error) {
	id, err := s.r.GetCafeIdsAndTotalByUserId(ctx, userId)
	return id, err
}

func NewService(r member.Requester) Service {
	return Service{r: r}
}
