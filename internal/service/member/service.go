package member

import (
	"cafe/internal/cli/member"
	req2 "cafe/internal/cli/member/req"
	"cafe/internal/page"
	"cafe/internal/service/member/req"
	"cafe/internal/service/member/res"
	"context"
	"errors"
)

type Service struct {
	r member.Requester
}

func NewService(r member.Requester) Service {
	return Service{r: r}
}

const (
	InvalidUserId = "invalid user id"
	InvalidCafeId = "invalid cafe id"
)

func (s Service) GetMemberInfo(ctx context.Context, cafeId int, userId int) (res.GetMemberInfo, error) {
	if userId == 0 {
		return res.GetMemberInfo{}, errors.New(InvalidUserId)
	}
	if cafeId == 0 {
		return res.GetMemberInfo{}, errors.New(InvalidCafeId)
	}

	md, err := s.r.GetCafeMyInfo(ctx, cafeId, userId)
	return res.GetMemberInfo{
		Id:        md.Id,
		UserId:    md.UserId,
		NickName:  md.NickName,
		CreatedAt: md.CreatedAt,
	}, err
}

func (s Service) GetCafeIdsAndTotal(ctx context.Context, userId int, reqPage page.ReqPage) ([]int, int, error) {
	id, total, err := s.r.GetCafeIdsAndTotalByUserId(ctx, userId, reqPage)
	return id, total, err
}

func (s Service) JoinCafe(ctx context.Context, j req.JoinCafe) error {
	err := s.r.JoinCafe(ctx, req2.JoinCafe{
		UserId:   j.UserId,
		CafeId:   j.CafeId,
		Nickname: j.Nickname,
	})
	return err
}

func (s Service) GetCafeMemberListCount(ctx context.Context, cafeId int, isBanned bool, reqPage page.ReqPage) ([]res.GetCafeMemberListCount, int, error) {
	listCunt, err := s.r.GetCafeMemberListCount(ctx, cafeId, isBanned, reqPage)
	if err != nil {
		return []res.GetCafeMemberListCount{}, 0, err
	}
	result := make([]res.GetCafeMemberListCount, len(listCunt.Members))
	for i, l := range listCunt.Members {
		result[i] = res.GetCafeMemberListCount{
			Id:        l.Id,
			UserId:    l.UserId,
			Nickname:  l.Nickname,
			CreatedAt: l.CreatedAt,
		}
	}
	return result, listCunt.Count, err
}

func (s Service) PatchMember(ctx context.Context, p req.PatchMember) error {
	err := s.r.PatchMember(ctx, req2.PatchMember{
		Nickname: p.Nickname,
		MemberId: p.MemberId,
	})
	return err
}

func (s Service) GetMemberInfoByCafeMemberId(ctx context.Context, cafeId int, memberId int) (res.GetMemberInfoByCafeMemberId, error) {
	if memberId == 0 {
		return res.GetMemberInfoByCafeMemberId{}, errors.New("invalid member id")
	}
	d, err := s.r.GetMemberByCafeMemberId(ctx, cafeId, memberId)
	if err != nil {
		return res.GetMemberInfoByCafeMemberId{}, err
	}
	return res.GetMemberInfoByCafeMemberId{
		Id:        d.Id,
		UserId:    d.UserId,
		Nickname:  d.NickName,
		CreatedAt: d.CreatedAt,
	}, err
}

func (s Service) GetMemberInfoByMemberIds(ctx context.Context, ids []int) ([]res.GetMemberInfoByMemberIds, error) {
	list, err := s.r.GetMemberListByMemberIds(ctx, ids)
	if err != nil {
		return []res.GetMemberInfoByMemberIds{}, err
	}
	result := make([]res.GetMemberInfoByMemberIds, len(list))
	for i, l := range list {
		result[i] = res.GetMemberInfoByMemberIds{
			Id:        l.Id,
			UserId:    l.UserId,
			Nickname:  l.Nickname,
			CreatedAt: l.CreatedAt,
		}
	}
	return result, err
}

func (s Service) CheckExistsByMemberId(ctx context.Context, cafeId, memberId int) (bool, error) {
	d, err := s.r.GetMemberByCafeMemberId(ctx, cafeId, memberId)
	if err != nil {
		return false, err
	}
	// 제로값 인지 확인
	return d.Id != 0, nil
}
