package memberRole

import (
	"cafe/internal/domain/memberRole/vo"
	"errors"
)

var _ MemberRole = (*memberRole)(nil)

type MemberRole interface {
	ValidCreate() error
	ValidUpdate() error

	ToInfo() vo.Info
	ToDetail() vo.Detail
}

type memberRole struct {
	id          int
	cafeId      int
	cafeRoleIds string
	memberId    int
}

const (
	InvalidMemberId = "invalid member id"
	InvalidId       = "invalid id"
)

func (m *memberRole) ValidCreate() error {
	if m.memberId < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (m *memberRole) ValidUpdate() error {
	if m.id < 1 {
		return errors.New(InvalidId)
	}
	if m.memberId < 1 {
		return errors.New(InvalidMemberId)
	}
	return nil
}

func (m *memberRole) ToInfo() vo.Info {
	return vo.Info{
		Id:          m.id,
		CafeRoleIds: m.cafeRoleIds,
	}
}

func (m *memberRole) ToDetail() vo.Detail {
	return vo.Detail{
		Id:          m.id,
		CafeRoleIds: m.cafeRoleIds,
		MemberId:    m.memberId,
	}
}
