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

func NewService(r member.Requester) Service {
	return Service{r: r}
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

func (s Service) GetCafeMemberListCount(ctx context.Context, cafeId int, isBanned bool, reqPage page.ReqPage) (domain.MemberListCount, error) {
	listCunt, err := s.r.GetCafeMemberListCount(ctx, cafeId, isBanned, reqPage)
	return listCunt, err
}

func (s Service) PatchMember(ctx context.Context, d domain.Member) error {
	err := s.r.PatchMember(ctx, d)
	return err
}

func (s Service) GetMemberInfoByCafeMemberId(ctx context.Context, cafeId int, memberId int) (domain.Member, error) {
	if memberId == 0 {
		return domain.Member{}, errors.New("invalid member id")
	}
	mDomain, err := s.r.GetMemberByCafeMemberId(ctx, cafeId, memberId)
	return mDomain, err
}

func (s Service) GetMemberInfoByMemberIds(ctx context.Context, ids []int) ([]domain.Member, error) {
	domains, err := s.r.GetMemberListByMemberIds(ctx, ids)
	return domains, err
}
