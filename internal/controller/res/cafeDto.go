package res

import "cafe/internal/domain"

type CafeListDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToListDtoList(cafes []domain.Cafe) []CafeListDto {
	dtos := make([]CafeListDto, len(cafes))
	for i, cafe := range cafes {
		dtos[i] = CafeListDto{
			Id:   cafe.Id,
			Name: cafe.Name,
		}
	}
	return dtos
}

type CafeDetailDto struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func ToDetailDto(cafe domain.Cafe) CafeDetailDto {
	return CafeDetailDto{
		Id:          cafe.Id,
		Name:        cafe.Name,
		Description: cafe.Description,
	}
}
