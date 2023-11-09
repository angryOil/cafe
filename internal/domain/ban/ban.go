package ban

import (
	"cafe/internal/domain/ban/vo"
	"errors"
	"time"
)

var _ Ban = (*ban)(nil)

type Ban interface {
	ValidCreate() error

	ToInfo() vo.Info
}

type ban struct {
	id          int
	userId      int
	memberId    int
	cafeId      int
	description string
	createdAt   time.Time
}

func (b *ban) ToInfo() vo.Info {
	return vo.Info{
		Id:          b.id,
		UserId:      b.userId,
		MemberId:    b.memberId,
		CafeId:      b.cafeId,
		Description: b.description,
		CreatedAt:   b.createdAt,
	}
}

const (
	InvalidCafeId   = "invalid cafe id"
	InvalidUserId   = "invalid user id"
	InvalidMemberId = "invalid member id"
)

func (b *ban) ValidCreate() error {
	if b.userId < 1 {
		return errors.New(InvalidUserId)
	}
	if b.memberId < 1 {
		return errors.New(InvalidMemberId)
	}
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	return nil
}
