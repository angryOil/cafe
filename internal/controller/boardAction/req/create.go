package req

type Create struct {
	ReadRoles   string `json:"read_roles"`
	CreateRoles string `json:"create_roles"`
}
