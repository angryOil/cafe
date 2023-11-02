package cDto

import "cafe/internal/domain"

type ListTotalDto[T any] struct {
	Contents []T `json:"contents"`
	Total    int `json:"total"`
}

func NewListTotalDto[T any](contents []T, total int) ListTotalDto[T] {
	return ListTotalDto[T]{
		Contents: contents,
		Total:    total,
	}
}

type MemberDetailRole struct {
	Id          int    `json:"id"`
	CafeRoleIds string `json:"cafe_role_ids"`
	MemberId    int    `json:"member_id"`
}

func ToDomainList(dtos []MemberDetailRole) []domain.MemberRole {
	results := make([]domain.MemberRole, len(dtos))
	for i, d := range dtos {
		results[i] = domain.MemberRole{
			Id:          d.Id,
			CafeRoleIds: d.CafeRoleIds,
			MemberId:    d.MemberId,
		}
	}
	return results
}

type MemberRole struct {
	Id          int    `json:"id"`
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (m MemberRole) ToDomain() domain.MemberRole {
	return domain.MemberRole{
		Id:          m.Id,
		CafeRoleIds: m.CafeRoleIds,
	}
}
