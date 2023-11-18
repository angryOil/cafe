package reply

import (
	"cafe/internal/cli/reply"
	"cafe/internal/service/reply/res"
	"context"
)

type Service struct {
	r reply.Requester
}

func NewService(r reply.Requester) Service {
	return Service{r: r}
}

func (s Service) GetList(ctx context.Context, boardId int) ([]res.GetList, int, error) {
	domains, total, err := s.r.GetList(ctx, boardId)
	if err != nil {
		return []res.GetList{}, 0, err
	}
	dto := make([]res.GetList, len(domains))
	for i, d := range domains {
		v := d.ToInfo()
		dto[i] = res.GetList{
			Id:            v.Id,
			Writer:        v.Writer,
			Content:       v.Content,
			CreatedAt:     v.CreatedAt,
			LastUpdatedAt: v.LastUpdatedAt,
		}
	}
	return dto, total, nil
}

func (s Service) GetCount(ctx context.Context, arr []int) ([]res.GetCount, error) {
	list, err := s.r.GetCount(ctx, arr)
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
