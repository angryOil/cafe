package board

import (
	"cafe/internal/domain/board/vo"
	"errors"
	"time"
)

var _ Board = (*board)(nil)

type Board interface {
	ValidCreate() error
	ValidUpdate() error

	ToListInfo() vo.ListInfo // 리스트 (간단 정보)
	ToDetail() vo.Detail     // 상세정보 (내용 포함)
	ToUpdate() vo.Update
}

type board struct {
	id            int
	cafeId        int
	boardType     int
	writer        int
	title         string
	content       string
	createdAt     string
	lastUpdatedAt string
}

const (
	InvalidBoardId   = "invalid board id"
	InvalidCafeId    = "invalid cafe id"
	InvalidBoardType = "invalid board type"
	InvalidWriterId  = "invalid writer id"
	InvalidTitle     = "invalid title"
	InvalidContent   = "invalid content"
)

func (b *board) ValidCreate() error {
	if b.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if b.title == "" {
		return errors.New(InvalidTitle)
	}
	if b.content == "" {
		return errors.New(InvalidContent)
	}
	if b.boardType < 1 {
		return errors.New(InvalidBoardType)
	}
	if b.writer < 1 {
		return errors.New(InvalidWriterId)
	}
	return nil
}

func (b *board) ValidUpdate() error {
	if b.id < 1 {
		return errors.New(InvalidContent)
	}
	if b.title == "" {
		return errors.New(InvalidTitle)
	}
	if b.content == "" {
		return errors.New(InvalidContent)
	}
	return nil
}

func (b *board) ToListInfo() vo.ListInfo {
	return vo.ListInfo{
		Id:        b.id,
		BoardType: b.boardType,
		Writer:    b.writer,
		Title:     b.title,
	}
}

func (b *board) ToDetail() vo.Detail {
	return vo.Detail{
		Id:        b.id,
		BoardType: b.boardType,
		Writer:    b.writer,
		Title:     b.title,
		Content:   b.content,
	}
}

func (b *board) ToUpdate() vo.Update {
	return vo.Update{
		Id:        b.id,
		CafeId:    b.cafeId,
		BoardType: b.boardType,
		Writer:    b.writer,
		Title:     b.title,
		Content:   b.content,
	}
}

var koreaZone, _ = time.LoadLocation("Asia/Seoul")

func convertTimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	t = t.In(koreaZone)
	return t.Format(time.RFC3339)
}
