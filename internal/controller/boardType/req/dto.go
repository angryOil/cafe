package req

import (
	"cafe/internal/domain"
)

type CreateBoardTypeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d CreateBoardTypeDto) ToDomain(cafeId, memberId int) domain.BoardType {
	return domain.BoardType{
		CreateBy:    memberId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	}
}

type PatchBoardTypeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d PatchBoardTypeDto) ToDomain(cafeId, typeId int) domain.BoardType {
	return domain.BoardType{
		Id:          typeId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	}
}
