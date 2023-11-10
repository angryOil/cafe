package req

type PutRole struct {
	CafeId      int
	MemberId    int
	CafeRoleIds string
}

type PutRoleDto struct {
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (p PutRole) ToDto() PutRoleDto {
	return PutRoleDto{CafeRoleIds: p.CafeRoleIds}
}
