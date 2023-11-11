package req

type Create struct {
	Name        string
	Description string
	CafeId      int
	OwnerId     int
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
