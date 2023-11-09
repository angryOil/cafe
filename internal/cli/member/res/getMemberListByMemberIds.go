package res

type GetMemberListByMemberIds struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Nickname  string `json:"nickname"`
	CreatedAt string `json:"created_at"`
}
