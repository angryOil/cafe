package reply

import (
	"cafe/internal/domain/reply/vo"
	"errors"
	"strings"
)

var _ Reply = (*reply)(nil)

type Reply interface {
	ValidCreate() error
	ValidUpdate() error

	ToInfo() vo.Info
	ToDetail() vo.Detail
}

type reply struct {
	id            int
	cafeId        int
	boardId       int
	writer        int
	content       string
	createdAt     string
	lastUpdatedAt string
}

const (
	InvalidId      = "invalid reply id"
	InvalidCafeId  = "invalid cafe id"
	InvalidBoardId = "invalid board id"
	InvalidContent = "invalid content"
)

func (r *reply) ValidCreate() error {
	if r.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if r.boardId < 1 {
		return errors.New(InvalidBoardId)
	}
	if strings.ReplaceAll(r.content, " ", "") == "" {
		return errors.New(InvalidContent)
	}
	return nil
}

func (r *reply) ValidUpdate() error {
	if r.id < 1 {
		return errors.New(InvalidId)
	}
	if strings.ReplaceAll(r.content, " ", "") == "" {
		return errors.New(InvalidContent)
	}
	return nil
}

func (r *reply) ToInfo() vo.Info {
	return vo.Info{
		Id:            r.id,
		Writer:        r.writer,
		Content:       r.content,
		CreatedAt:     r.createdAt,
		LastUpdatedAt: r.lastUpdatedAt,
	}
}

func (r *reply) ToDetail() vo.Detail {
	return vo.Detail{
		Id:            r.id,
		CafeId:        r.cafeId,
		BoardId:       r.boardId,
		Writer:        r.writer,
		Content:       r.content,
		CreatedAt:     r.createdAt,
		LastUpdatedAt: r.lastUpdatedAt,
	}
}
