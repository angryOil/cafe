package res

type Info struct {
	Id            int    `json:"id,omitempty"`
	BoardType     int    `json:"board_type_id,omitempty"`
	Writer        int    `json:"writer_id,omitempty"`
	Title         string `json:"title,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"lastUpdated_at,omitempty"`
}

type ListTotalDto struct {
	Content []Info `json:"content"`
	Total   int    `json:"total"`
}
