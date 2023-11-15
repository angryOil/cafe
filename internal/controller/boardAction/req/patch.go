package req

type Patch struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
	UpdateRoles string `json:"update_roles"`
	UpdateAble  bool   `json:"update_able"`
	DeleteRoles string `json:"delete_roles"`
}
