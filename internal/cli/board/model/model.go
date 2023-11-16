package model

import "cafe/internal/domain/board"

type Board struct {
	Id            int    `json:"id,omitempty"`
	BoardType     int    `json:"board_type_id,omitempty"`
	Writer        int    `json:"writer_id,omitempty"`
	Title         string `json:"title,omitempty"`
	Content       string `json:"content,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	LastUpdatedAt string `json:"lastUpdated_at,omitempty"`
}

func (b Board) ToDomain() board.Board {
	return board.NewBuilder().
		Id(b.Id).
		BoardType(b.BoardType).
		Writer(b.Writer).
		Title(b.Title).
		Content(b.Content).
		CreatedAt(b.CreatedAt).
		LastUpdatedAt(b.LastUpdatedAt).
		Build()
}

func ToDomainList(bList []Board) []board.Board {
	result := make([]board.Board, len(bList))
	for i, b := range bList {
		result[i] = b.ToDomain()
	}
	return result
}

type BoardListTotalDto struct {
	Content []Board `json:"content"`
	Total   int     `json:"total"`
}
