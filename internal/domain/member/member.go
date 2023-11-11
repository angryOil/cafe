package member

import (
	"cafe/internal/domain/member/vo"
	"errors"
	"time"
)

var _ Member = (*member)(nil)

type Member interface {
	ValidCreate() error
	ValidUpdate() error

	Update(nickname string) Member

	ToInfo() vo.Info
}

type member struct {
	id        int
	cafeId    int
	userId    int
	nickname  string
	createdAt time.Time
}

const (
	InvalidNickname = "invalid nickname"
	InvalidUserId   = "invalid user id"
	InvalidCafeId   = "invalid cafe id"
	InvalidId       = "invalid id"
)

func (m *member) ValidCreate() error {
	if m.nickname == "" {
		return errors.New(InvalidNickname)
	}
	if m.userId < 1 {
		return errors.New(InvalidUserId)
	}
	if m.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	return nil
}

func (m *member) ValidUpdate() error {
	if m.nickname == "" {
		return errors.New(InvalidNickname)
	}
	if m.userId < 1 {
		return errors.New(InvalidUserId)
	}
	if m.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if m.id < 1 {
		return errors.New(InvalidId)
	}
	return nil
}

func (m *member) Update(nickname string) Member {
	m.nickname = nickname
	return m
}

func (m *member) ToInfo() vo.Info {
	return vo.Info{
		Id:        m.id,
		UserId:    m.userId,
		Nickname:  m.nickname,
		CreatedAt: m.createdAt,
	}
}
