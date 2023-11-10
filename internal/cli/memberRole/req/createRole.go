package req

type CreateRole struct {
	CafeId      int
	MemberId    int
	CafeRoleIds string
}

type CreateRoleDto struct {
	CafeRoleIds string `json:"cafe_role_ids"`
}

func (c CreateRole) ToDto() CreateRoleDto {
	return CreateRoleDto{CafeRoleIds: c.CafeRoleIds}
}
