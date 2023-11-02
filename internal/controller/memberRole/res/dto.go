package res

import "cafe/internal/domain"

type MemberRoleDto struct {
	Id          int    `json:"id"`
	CafeRoleIds string `json:"cafe_role_ids"`
}

type RoleArrDto struct {
	Id    int    `json:"id"`
	Roles []Role `json:"roles"`
}

func (m MemberRoleDto) ToRoleArrDto(roles []Role) RoleArrDto {
	return RoleArrDto{
		Id:    m.Id,
		Roles: roles,
	}
}

func ToMemberRoleDto(d domain.MemberRole) MemberRoleDto {
	return MemberRoleDto{
		Id:          d.Id,
		CafeRoleIds: d.CafeRoleIds,
	}
}

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
