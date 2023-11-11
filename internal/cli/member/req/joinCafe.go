package req

type JoinCafe struct {
	UserId   int
	CafeId   int
	Nickname string
}

type JoinDto struct {
	Nickname string `json:"nickname"`
}

func (j JoinCafe) ToJoinDto() JoinDto {
	return JoinDto{Nickname: j.Nickname}
}
