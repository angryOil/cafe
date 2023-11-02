package req

import (
	"cafe/internal/domain"
)

type PutMemberRoleDto struct {
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (d PutMemberRoleDto) ToDomain() domain.MemberRole {
	return domain.MemberRole{
		CafeRoleIds: d.CafeRoleIds,
	}
}
