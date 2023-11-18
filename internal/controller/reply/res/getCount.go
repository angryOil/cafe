package res

type GetCount struct {
	BoardId    int `json:"board_id"`
	ReplyCount int `json:"reply_count"`
}

type CountListDto struct {
	Content []GetCount `json:"content"`
}
