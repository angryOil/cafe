package req

type Patch struct {
	Id          int
	CafeId      int
	BoardTypeId int
	ReadRoles   string
	CreateRoles string
	UpdateRoles string
	UpdateAble  bool
	DeleteRoles string
}

func (p Patch) ToPatchDto() PatchDto {
	return PatchDto{
		ReadRoles:   p.ReadRoles,
		CreateRoles: p.CreateRoles,
		UpdateRoles: p.UpdateRoles,
		UpdateAble:  p.UpdateAble,
		DeleteRoles: p.DeleteRoles,
	}
}

type PatchDto struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
	UpdateRoles string `json:"update_roles"`
	UpdateAble  bool   `json:"update_able"`
	DeleteRoles string `json:"delete_roles"`
}
