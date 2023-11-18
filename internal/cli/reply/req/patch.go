package req

type Patch struct {
	Id      int
	Content string
}

func (p Patch) ToPatchDto() PatchDto {
	return PatchDto{Content: p.Content}
}

type PatchDto struct {
	Content string `json:"content"`
}
