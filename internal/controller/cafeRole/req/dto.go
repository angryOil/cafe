package req

import "cafe/internal/domain"

type CreateRoleDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d CreateRoleDto) ToDomain(cafeId int) domain.CafeRole {
	return domain.CafeRole{
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	}
}

type PatchRoleDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (d PatchRoleDto) ToDomain(cafeId, roleId int) domain.CafeRole {
	return domain.CafeRole{
		CafeId:      cafeId,
		Id:          roleId,
		Name:        d.Name,
		Description: d.Description,
	}
}
