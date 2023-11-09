package req

type PatchMember struct {
	Nickname string
	MemberId int
}

type PatchDto struct {
	Nickname string `json:"nickname,omitempty"`
}

func (p PatchMember) ToPatchDto() PatchDto {
	return PatchDto{Nickname: p.Nickname}
}
