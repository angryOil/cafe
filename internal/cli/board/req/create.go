package req

type Create struct {
	Writer    int
	CafeId    int
	BoardType int
	Title     string
	Content   string
}

func (c Create) ToCreateDto() CreateDto {
	return CreateDto{
		Writer:  c.Writer,
		Title:   c.Title,
		Content: c.Content,
	}
}

type CreateDto struct {
	Writer  int    `json:"writer_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
