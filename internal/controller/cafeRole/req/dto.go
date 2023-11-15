package req

type CreateRoleDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type PatchRoleDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
