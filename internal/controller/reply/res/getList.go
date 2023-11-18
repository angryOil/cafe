package res

type GetList struct {
	Id            int    `json:"id,omitempty"`
	Writer        int    `json:"writer_id,omitempty"`
	Content       string `json:"content,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"last_updated_at,omitempty"`
}

type ListTotalDto struct {
	Content []GetList `json:"content"`
	Total   int       `json:"total"`
}
