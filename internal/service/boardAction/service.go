package boardAction

import (
	"cafe/internal/cli/boardAction"
	"cafe/internal/service/boardAction/res"
	"context"
)

type Service struct {
	r boardAction.Requester
}

func NewService(r boardAction.Requester) Service {
	return Service{r: r}
}

func (s Service) GetInfo(ctx context.Context, cafeId, boardTypeId int) (res.GetInfo, error) {
	d, err := s.r.GetInfo(ctx, cafeId, boardTypeId)
	if err != nil {
		return res.GetInfo{}, err
	}
	v := d.ToInfo()
	return res.GetInfo{
		Id:          v.Id,
		CafeId:      v.CafeId,
		BoardTypeId: v.BoardTypeId,
		ReadRoles:   v.ReadRoles,
		CreateRoles: v.CreateRoles,
		UpdateRoles: v.UpdateRoles,
		UpdateAble:  v.UpdateAble,
		DeleteRoles: v.DeleteRoles,
	}, nil
}
