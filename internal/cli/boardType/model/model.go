package model

import "cafe/internal/domain/boardType"

type BoardType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BoardTypeListCount struct {
	BoardTypes []BoardType `json:"contents"`
	Total      int         `json:"total"`
}

func (b BoardType) ToDomain() boardType.BoardType {
	return boardType.NewBuilder().
		Id(b.Id).
		Name(b.Name).
		Description(b.Description).
		Build()
}
