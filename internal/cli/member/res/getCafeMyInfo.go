package res

type GetCafeMyInfo struct {
	Id        int    `json:"member_id,omitempty"`
	UserId    int    `json:"user_id,omitempty"`
	NickName  string `json:"nickname,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
