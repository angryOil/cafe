package req

type Patch struct {
	Id          int
	CafeId      int
	Name        string
	Description string
}

func (p Patch) ToPatchDto() PatchDto {
	return PatchDto{
		Name:        p.Name,
		Description: p.Description,
	}
}

type PatchDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
