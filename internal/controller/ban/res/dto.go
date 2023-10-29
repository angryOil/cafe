package res

import "cafe/internal/domain"

type BanListDto struct {
	Id          int    `json:"id"`
	CafeId      int    `json:"cafe_id"`
	Description string `json:"description"`
}

func ToBandListDtoList(domains []domain.Ban) []BanListDto {
	results := make([]BanListDto, len(domains))
	for i, d := range domains {
		results[i] = BanListDto{
			Id:          d.Id,
			CafeId:      d.CafeId,
			Description: d.Description,
		}
	}
	return results
}
