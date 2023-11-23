package res

type Detail struct {
	Id            int     `json:"id,omitempty"`
	BoardType     int     `json:"board_type_id,omitempty"`
	Writer        int     `json:"writer_id,omitempty"`
	Title         string  `json:"title,omitempty"`
	Content       string  `json:"content,omitempty"`
	CreatedAt     string  `json:"created_at,omitempty"`
	LastUpdatedAt string  `json:"lastUpdated_at,omitempty"`
	Replies       []Reply `json:"replies,omitempty"`
	ReplyCnt      int     `json:"reply_cnt,omitempty"`
}

func (d *Detail) SetReplies(replies []Reply, replyCnt int) {
	d.Replies = replies
	d.ReplyCnt = replyCnt
}

type Reply struct {
	Id            int    `json:"id,omitempty"`
	Writer        int    `json:"writer_id,omitempty"`
	Content       string `json:"content,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"last_updated_at,omitempty"`
}
