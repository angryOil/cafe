package model

import "cafe/internal/domain/reply"

type Reply struct {
	Id            int    `json:"id,omitempty"`
	Writer        int    `json:"writer_id,omitempty"`
	Content       string `json:"content,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"last_updated_at,omitempty"`
}

type ListTotalDto struct {
	Content []Reply `json:"content"`
	Total   int     `json:"total"`
}

func ToDomainList(rList []Reply) []reply.Reply {
	result := make([]reply.Reply, len(rList))
	for i, r := range rList {
		result[i] = reply.NewBuilder().
			Id(r.Id).
			Writer(r.Writer).
			Content(r.Content).
			CreatedAt(r.CreatedAt).
			LastUpdatedAt(r.LastUpdatedAt).
			Build()
	}
	return result
}
