package model

type BoardType struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BoardTypeListCount struct {
	BoardTypes []BoardType `json:"contents"`
	Total      int         `json:"total"`
}
