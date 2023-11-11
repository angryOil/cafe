package req

type PutRole struct {
	Id          int
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
