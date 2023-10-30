package res

import "cafe/internal/domain"

type BoardTypeDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToBoardTypeDtoList(domains []domain.BoardType) []BoardTypeDto {
	results := make([]BoardTypeDto, len(domains))
	for i, d := range domains {
		results[i] = BoardTypeDto{
			Id:          d.Id,
			Name:        d.Name,
			Description: d.Description,
		}
	}
	return results
}
