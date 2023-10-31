package dto

import "cafe/internal/domain"

type RoleListTotalCDto struct {
	Roles []RoleCDto `json:"contents"`
	Total int        `json:"total"`
}

type RoleCDto struct {
	Id          int
	Name        string
	Description string
}

func ToDomainList(dtos []RoleCDto) []domain.CafeRole {
	results := make([]domain.CafeRole, len(dtos))
	for i, d := range dtos {
		results[i] = domain.CafeRole{
			Id:          d.Id,
			Name:        d.Name,
			Description: d.Description,
		}
	}
	return results
}

type CreateRoleCDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCreateRoleCDto(d domain.CafeRole) CreateRoleCDto {
	return CreateRoleCDto{
		Name:        d.Name,
		Description: d.Description,
	}
}

type PatchRoleCDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToPatchRoleCDto(d domain.CafeRole) PatchRoleCDto {
	return PatchRoleCDto{
		Name:        d.Name,
		Description: d.Description,
	}
}
