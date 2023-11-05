package res

import (
	"cafe/internal/domain/cafe_vo"
)

type CafeList struct {
	Contents []CafeListDto `json:"contents"`
}

type CafeListDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToListDtoList(pageInfos []cafe_vo.CafeListInfo) []CafeListDto {
	dto := make([]CafeListDto, len(pageInfos))
	for i, info := range pageInfos {
		dto[i] = CafeListDto{
			Id:   info.Id,
			Name: info.Name,
		}
	}
	return dto
}

type CafeDetailDto struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func ToDetailDto(detail cafe_vo.Detail) CafeDetailDto {
	return CafeDetailDto{
		Id:          detail.Id,
		Name:        detail.Name,
		Description: detail.Description,
	}
}
