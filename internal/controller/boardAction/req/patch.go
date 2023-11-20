package req

type Patch struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
}
