package res

import "cafe/internal/domain"

type DetailDto struct {
	Id          int    `json:"id"`
	CafeRoleIds string `json:"cafe_role_ids"`
	MemberId    int    `json:"member_id"`
}

func ToDetailList(domains []domain.MemberRole) []DetailDto {
	results := make([]DetailDto, len(domains))
	for i, d := range domains {
		results[i] = DetailDto{
			Id:          d.Id,
			CafeRoleIds: d.CafeRoleIds,
			MemberId:    d.MemberId,
		}
	}
	return results
}

type Role struct {
	RoleId int    `json:"role_id"`
	Name   string `json:"name"`
}

type DetailRoleArrDto struct {
	Id       int    `json:"id"`
	MemberID int    `json:"member_id"`
	Roles    []Role `json:"roles"`
}

func (detail DetailDto) ToRoleArrDto(roles []Role) DetailRoleArrDto {
	return DetailRoleArrDto{
		Id:       detail.Id,
		MemberID: detail.MemberId,
		Roles:    roles,
	}
}
