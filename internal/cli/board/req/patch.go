package req

type Patch struct {
	Id        int
	Requester int
	Title     string
	Content   string
}

func (p Patch) ToPatchDto() PatchDto {
	return PatchDto{
		Requester: p.Requester,
		Title:     p.Title,
		Content:   p.Content,
	}
}

type PatchDto struct {
	Requester int    `json:"requester_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
