package res

type MemberInfoDto struct {
	Id        int    `json:"member_id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type IdsTotalDto struct {
	Ids   []int `json:"ids"`
	Total int   `json:"total"`
}
