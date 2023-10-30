package dto

import (
	"cafe/internal/domain"
)

type BoardTypeCDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// 사실 모두 같은 방식으로 count,contents 를 줄거지만 서비스마다 요구가 달라질수있다고
// 산정하고 따로만듬

type BoardTypeCDtoListCount struct {
	Boards []BoardTypeCDto `json:"contents"`
	Total  int             `json:"total"`
}

func ToDomainList(list []BoardTypeCDto) []domain.BoardType {
	results := make([]domain.BoardType, len(list))
	for i, cDto := range list {
		results[i] = domain.BoardType{
			Id:          cDto.Id,
			Name:        cDto.Name,
			Description: cDto.Description,
		}
	}
	return results
}

type CreateBoardTypeCDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToCreateBoardTypeCDto(d domain.BoardType) CreateBoardTypeCDto {
	return CreateBoardTypeCDto{
		Name:        d.Name,
		Description: d.Description,
	}
}

type PatchBoardTypeCDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToPatchBoardTypeCDto(d domain.BoardType) CreateBoardTypeCDto {
	return CreateBoardTypeCDto{
		Name:        d.Name,
		Description: d.Description,
	}
}
