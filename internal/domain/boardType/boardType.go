package boardType

import (
	"cafe/internal/domain/boardType/vo"
	"errors"
	"time"
)

var _ BoardType = (*boardType)(nil)

type BoardType interface {
	ValidCreate() error
	ValidUpdate() error
	ToInfo() vo.Info
}

const (
	InvalidCreateBy    = "invalid created by"
	InvalidCafeId      = "invalid cafe id"
	InvalidName        = "invalid name"
	InvalidBoardTypeId = "invalid board type id"
)

type boardType struct {
	id          int
	createBy    int
	cafeId      int
	name        string
	description string
	createdAt   time.Time
}

func (b *boardType) ValidUpdate() error {
	if b.id < 1 {
		return errors.New(InvalidBoardTypeId)
	}
	if b.name == "" {
		return errors.New(InvalidName)
	}
	return nil
}

func (b *boardType) ToInfo() vo.Info {
	return vo.Info{
		Id:          b.id,
		Name:        b.name,
		Description: b.description,
	}
}

func (b *boardType) ValidCreate() error {
	if b.createBy < 1 {
		return errors.New(InvalidCreateBy)
	}
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.name == "" {
		return errors.New(InvalidName)
	}
	return nil
}
