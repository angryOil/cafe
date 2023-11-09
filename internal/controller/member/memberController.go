package member

import (
	"cafe/internal/controller/member/req"
	"cafe/internal/controller/member/res"
	"cafe/internal/page"
	"cafe/internal/service/member"
	req2 "cafe/internal/service/member/req"
	"context"
)

type Controller struct {
	s member.Service
}

func NewController(s member.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetMemberInfo(ctx context.Context, cafeId int, userId int) (res.MemberInfoDto, error) {
	md, err := c.s.GetMemberInfo(ctx, cafeId, userId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	return res.MemberInfoDto{
		Id:        md.Id,
		UserId:    md.UserId,
		Nickname:  md.NickName,
		CreatedAt: md.CreatedAt,
	}, err
}

func (c Controller) GetMyCafeIds(ctx context.Context, userId int, reqPage page.ReqPage) ([]int, int, error) {
	ids, total, err := c.s.GetCafeIdsAndTotal(ctx, userId, reqPage)
	if err != nil {
		return []int{}, 0, err
	}
	return ids, total, nil
}

func (c Controller) JoinCafe(ctx context.Context, userId, cafeId int, dto req.JoinMemberDto) error {
	err := c.s.JoinCafe(ctx, req2.JoinCafe{
		CafeId:   cafeId,
		UserId:   userId,
		Nickname: dto.Nickname,
	})
	return err
}

func (c Controller) GetCafeMemberListCount(ctx context.Context, cafeId int, isBanned bool, reqPage page.ReqPage) ([]res.MemberInfoDto, int, error) {
	list, cnt, err := c.s.GetCafeMemberListCount(ctx, cafeId, isBanned, reqPage)
	if err != nil {
		return []res.MemberInfoDto{}, 0, err
	}
	result := make([]res.MemberInfoDto, len(list))
	for i, l := range list {
		result[i] = res.MemberInfoDto{
			Id:        l.Id,
			UserId:    l.UserId,
			Nickname:  l.Nickname,
			CreatedAt: l.CreatedAt,
		}
	}
	return result, cnt, nil
}

func (c Controller) PatchMember(ctx context.Context, memberId int, dto req.PatchMemberDto) error {
	err := c.s.PatchMember(ctx, req2.PatchMember{
		Nickname: dto.Nickname,
		MemberId: memberId,
	})
	return err
}

func (c Controller) GetInfoByCafeMemberId(ctx context.Context, cafeId int, memberId int) (res.MemberInfoDto, error) {
	d, err := c.s.GetMemberInfoByCafeMemberId(ctx, cafeId, memberId)
	if err != nil {
		return res.MemberInfoDto{}, err
	}
	return res.MemberInfoDto{
		Id:        d.Id,
		UserId:    d.UserId,
		Nickname:  d.Nickname,
		CreatedAt: d.CreatedAt,
	}, nil
}

func (c Controller) GetMembersByMemberIds(ctx context.Context, memberIds []int) ([]res.MemberInfoDto, error) {
	list, err := c.s.GetMemberInfoByMemberIds(ctx, memberIds)
	if err != nil {
		return []res.MemberInfoDto{}, err
	}
	result := make([]res.MemberInfoDto, len(list))
	for i, l := range list {
		result[i] = res.MemberInfoDto{
			Id:        l.Id,
			UserId:    l.UserId,
			Nickname:  l.Nickname,
			CreatedAt: l.CreatedAt,
		}
	}
	return result, err
}

func (c Controller) CheckExistsMemberByMemberId(ctx context.Context, cafeId, memberId int) (bool, error) {
	ok, err := c.s.CheckExistsByMemberId(ctx, cafeId, memberId)
	return ok, err
}
