package reply

import (
	"cafe/internal/cli/reply"
	"cafe/internal/cli/reply/req"
	reply2 "cafe/internal/domain/reply"
	req2 "cafe/internal/service/reply/req"
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

func (s Service) Create(ctx context.Context, c req2.Create) error {
	writer, boardId, cafeId := c.Writer, c.BoardId, c.CafeId
	content := c.Content
	err := reply2.NewBuilder().
		CafeId(cafeId).
		BoardId(boardId).
		Writer(writer).
		Content(content).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.r.Create(ctx, req.Create{
		BoardId: boardId,
		CafeId:  cafeId,
		Writer:  writer,
		Content: content,
	})
	return err
}

func (s Service) Delete(ctx context.Context, replyId int) error {
	err := s.r.Delete(ctx, replyId)
	return err
}

func (s Service) Patch(ctx context.Context, p req2.Patch) error {
	id := p.Id
	content := p.Content
	err := reply2.NewBuilder().
		Id(id).
		Content(content).
		Build().ValidUpdate()
	if err != nil {
		return err
	}

	err = s.r.Patch(ctx, req.Patch{
		Id:      id,
		Content: content,
	})
	return err
}
