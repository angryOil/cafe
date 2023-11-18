package reply

import (
	"cafe/internal/controller/reply/req"
	"cafe/internal/controller/reply/res"
	"cafe/internal/service/reply"
	req2 "cafe/internal/service/reply/req"
	"context"
)

type Controller struct {
	s reply.Service
}

func NewController(s reply.Service) Controller {
	return Controller{s: s}
}

func (c Controller) GetList(ctx context.Context, boardId int) ([]res.GetList, int, error) {
	list, total, err := c.s.GetList(ctx, boardId)
	if err != nil {
		return []res.GetList{}, 0, err
	}
	dto := make([]res.GetList, len(list))
	for i, l := range list {
		dto[i] = res.GetList{
			Id:            l.Id,
			Writer:        l.Writer,
			Content:       l.Content,
			CreatedAt:     l.CreatedAt,
			LastUpdatedAt: l.LastUpdatedAt,
		}
	}
	return dto, total, nil
}

func (c Controller) GetCount(ctx context.Context, arr []int) ([]res.GetCount, error) {
	list, err := c.s.GetCount(ctx, arr)
	if err != nil {
		return []res.GetCount{}, err
	}
	dto := make([]res.GetCount, len(list))
	for i, l := range list {
		dto[i] = res.GetCount{
			BoardId:    l.BoardId,
			ReplyCount: l.ReplyCount,
		}
	}
	return dto, nil
}

func (c Controller) Create(ctx context.Context, cafeId int, boardId int, memberId int, create req.Create) error {
	err := c.s.Create(ctx, req2.Create{
		BoardId: boardId,
		CafeId:  cafeId,
		Writer:  memberId,
		Content: create.Content,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, replyId int) error {
	err := c.s.Delete(ctx, replyId)
	return err
}

func (c Controller) Patch(ctx context.Context, replyId int, p req.Patch) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:      replyId,
		Content: p.Content,
	})
	return err
}
