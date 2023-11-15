package boardAction

import (
	"cafe/internal/domain/boardAction/vo"
	"errors"
)

var _ BoardAction = (*boardAction)(nil)

type BoardAction interface {
	ValidCreate() error
	ValidUpdate() error

	ToInfo() vo.Info
}

type boardAction struct {
	id          int
	cafeId      int
	boardTypeId int
	readRoles   string
	createRoles string
	updateRoles string
	updateAble  bool
	deleteRoles string
}

const (
	InvalidId          = "invalid board action id"
	InvalidCafeId      = "invalid cafe id"
	InvalidBoardTypeId = "invalid board type id"
)

func (b *boardAction) ValidCreate() error {
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.boardTypeId < 1 {
		return errors.New(InvalidBoardTypeId)
	}
	return nil
}

func (b *boardAction) ValidUpdate() error {
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.boardTypeId < 1 {
		return errors.New(InvalidBoardTypeId)
	}
	if b.id < 1 {
		return errors.New(InvalidId)
	}
	return nil
}

func (b *boardAction) ToInfo() vo.Info {
	return vo.Info{
		Id:          b.id,
		CafeId:      b.cafeId,
		BoardTypeId: b.boardTypeId,
		ReadRoles:   b.readRoles,
		CreateRoles: b.createRoles,
		UpdateRoles: b.updateRoles,
		UpdateAble:  b.updateAble,
		DeleteRoles: b.deleteRoles,
	}
}
