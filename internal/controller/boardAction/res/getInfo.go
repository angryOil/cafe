package res

type GetInfo struct {
	Id          int    `json:"id"`
	CafeId      int    `json:"cafe_id"`
	BoardTypeId int    `json:"board_type_id"`
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
	UpdateRoles string `json:"update_roles"`
	UpdateAble  bool   `json:"update_able"`
	DeleteRoles string `json:"delete_roles"`
}
