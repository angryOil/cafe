package res

type Detail struct {
	Id            int    `json:"id,omitempty"`
	BoardType     int    `json:"board_type_id,omitempty"`
	Writer        int    `json:"writer_id,omitempty"`
	Title         string `json:"title,omitempty"`
	Content       string `json:"content,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"lastUpdated_at,omitempty"`
}
