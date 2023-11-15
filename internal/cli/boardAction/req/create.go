package req

type Create struct {
	CafeId      int
	BoardType   int
	ReadRoles   string
	CreateRoles string
	UpdateRoles string
	UpdateAble  bool
	DeleteRoles string
}

type CreateDto struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
	UpdateRoles string `json:"update_roles"`
	UpdateAble  bool   `json:"update_able"`
	DeleteRoles string `json:"delete_roles"`
}

func (c Create) ToCreateDto() CreateDto {
	return CreateDto{
		ReadRoles:   c.ReadRoles,
		CreateRoles: c.CreateRoles,
		UpdateRoles: c.UpdateRoles,
		UpdateAble:  c.UpdateAble,
		DeleteRoles: c.DeleteRoles,
	}
}
