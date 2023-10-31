package res

import "cafe/internal/domain"

type RoleDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToDtoList(domains []domain.CafeRole) []RoleDto {
	results := make([]RoleDto, len(domains))
	for i, d := range domains {
		results[i] = RoleDto{
			Id:          d.Id,
			Name:        d.Name,
			Description: d.Description,
		}
	}
	return results
}
