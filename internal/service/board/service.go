package board

import (
	"cafe/internal/cli/board"
	req2 "cafe/internal/cli/board/req"
	board2 "cafe/internal/domain/board"
	page2 "cafe/internal/page"
	"cafe/internal/service/board/req"
	"cafe/internal/service/board/res"
	"context"
)

type Service struct {
	r board.Requester
}

func NewService(r board.Requester) Service {
	return Service{r: r}
}

func (s Service) GetList(ctx context.Context, l req.GetList, reqPage page2.ReqPage) ([]res.GetList, int, error) {
	domains, total, err := s.r.GetList(ctx, req2.GetList{
		CafeId:    l.CafeId,
		BoardType: l.BoardType,
		Writer:    l.Writer,
	}, reqPage)
	if err != nil {
		return []res.GetList{}, 0, err
	}
	dto := make([]res.GetList, len(domains))
	for i, d := range domains {
		v := d.ToListInfo()
		dto[i] = res.GetList{
			Id:            v.Id,
			BoardType:     v.BoardType,
			Writer:        v.Writer,
			Title:         v.Title,
			CreatedAt:     v.CreatedAt,
			LastUpdatedAt: v.LastUpdatedAt,
		}
	}
	return dto, total, nil
}

func (s Service) Create(ctx context.Context, c req.Create) error {
	err := board2.NewBuilder().
		Writer(c.Writer).
		CafeId(c.CafeId).
		BoardType(c.BoardType).
		Title(c.Title).
		Content(c.Content).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.r.Create(ctx, req2.Create{
		Writer:    c.Writer,
		CafeId:    c.CafeId,
		BoardType: c.BoardType,
		Title:     c.Title,
		Content:   c.Content,
	})
	return err
}

func (s Service) Patch(ctx context.Context, p req.Patch) error {
	err := board2.NewBuilder().
		Id(p.Id).
		Title(p.Title).
		Content(p.Content).
		Build().ValidUpdate()
	if err != nil {
		return err
	}

	err = s.r.Patch(ctx, req2.Patch{
		Id:        p.Id,
		Title:     p.Title,
		Content:   p.Content,
		Requester: p.Requester,
	})
	return err
}

func (s Service) Delete(ctx context.Context, id int) error {
	err := s.r.Delete(ctx, id)
	return err
}
