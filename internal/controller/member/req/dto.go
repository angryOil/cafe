package req

type JoinMemberDto struct {
	Nickname string `json:"nickname"`
}

type PatchMemberDto struct {
	Nickname string `json:"nickname"`
}
