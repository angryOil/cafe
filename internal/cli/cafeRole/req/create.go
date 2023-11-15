package req

type Create struct {
	CafeId      int
	Name        string
	Description string
}

func (c Create) ToCreateDto() CreateDto {
	return CreateDto{
		Name:        c.Name,
		Description: c.Description,
	}
}

type CreateDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
