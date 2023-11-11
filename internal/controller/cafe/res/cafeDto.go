package res

import (
	cafe_vo2 "cafe/internal/domain/cafe/vo"
)

type CafeList struct {
	Contents []CafeListDto `json:"contents"`
}

type CafeListDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToListDtoList(pageInfos []cafe_vo2.CafeListInfo) []CafeListDto {
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

func ToDetailDto(detail cafe_vo2.Detail) CafeDetailDto {
	return CafeDetailDto{
		Id:          detail.Id,
		Name:        detail.Name,
		Description: detail.Description,
	}
}
