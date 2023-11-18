package req

type Create struct {
	BoardId int
	CafeId  int
	Writer  int
	Content string
}

func (c Create) ToCreateDto() CreateDto {
	return CreateDto{
		Writer:  c.Writer,
		Content: c.Content,
	}
}

type CreateDto struct {
	Writer  int    `json:"writer_id"`
	Content string `json:"content"`
}
