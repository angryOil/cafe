package model

import cafeRole2 "cafe/internal/domain/cafeRole"

type CafeRole struct {
	Id          int    `json:"id"`
	CafeId      int    `json:"cafeId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CafeRoleListTotal struct {
	Roles []CafeRole `json:"contents"`
	Total int        `json:"total"`
}

func (c CafeRole) ToDomain() cafeRole2.CafeRole {
	return cafeRole2.NewBuilder().
		Id(c.Id).
		CafeId(c.CafeId).
		Name(c.Name).
		Description(c.Description).
		Build()
}
