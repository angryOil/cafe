package reply

import (
	"cafe/internal/controller/reply/res"
	"cafe/internal/service/reply"
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
