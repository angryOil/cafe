package model

import "cafe/internal/domain/memberRole"

type MemberRole struct {
	Id          int    `json:"id"`
	CafeRoleIds string `json:"cafe_role_ids"`
	MemberId    int    `json:"member_id"`
}

func (m MemberRole) ToDomain() memberRole.MemberRole {
	return memberRole.NewBuilder().
		Id(m.Id).
		CafeRoleIds(m.CafeRoleIds).
		MemberId(m.MemberId).
		Build()
}

func ToDomainList(models []MemberRole) []memberRole.MemberRole {
	result := make([]memberRole.MemberRole, len(models))
	for i, m := range models {
		result[i] = m.ToDomain()
	}
	return result
}

type ListTotalDto struct {
	Contents []MemberRole `json:"contents"`
	Total    int          `json:"total"`
}
