package model

import (
	domain "cafe/internal/domain"
	"github.com/uptrace/bun"
	"time"
)

type Cafe struct {
	bun.BaseModel `bun:"table:cafe,alias:t"`

	Id          int       `bun:"id,pk,autoincrement"`
	OwnerId     int       `bun:"owner_id,notnull"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,notnull"`
}

func ToModel(cd domain.Cafe) Cafe {
	return Cafe{
		Id:          cd.Id,
		OwnerId:     cd.OwnerId,
		Name:        cd.Name,
		Description: cd.Description,
		CreatedAt:   cd.CreatedAt,
	}
}

func ToDomainDetailList(list []Cafe) []domain.Cafe {
	results := make([]domain.Cafe, len(list))
	for i, detail := range list {
		results[i] = domain.Cafe{
			Id:          detail.Id,
			OwnerId:     detail.OwnerId,
			Name:        detail.Name,
			Description: detail.Description,
			CreatedAt:   detail.CreatedAt,
		}
	}
	return results
}

type CateList struct {
	bun.BaseModel `bun:"table:cafe,alias:t"`

	Id          int    `bun:"id,pk,autoincrement"`
	Name        string `bun:"name,notnull"`
	Description string `bun:"description"`
}

func ToDomainList(list []CateList) []domain.Cafe {
	results := make([]domain.Cafe, len(list))
	for i, cafe := range list {
		results[i] = domain.Cafe{
			Id:          cafe.Id,
			Name:        cafe.Name,
			Description: cafe.Description,
		}
	}
	return results
}
