package model

import (
	domain "cafe/internal/domain"
	"cafe/internal/repository/request"
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

func ToCreateModel(cd request.CreateCafe) Cafe {
	return Cafe{
		OwnerId:     cd.OwnerId,
		Name:        cd.Name,
		Description: cd.Description,
		CreatedAt:   cd.CreatedAt,
	}
}

func ToSaveModel(uc request.UpdateCafe) Cafe {
	return Cafe{
		Name:        uc.Name,
		Description: uc.Description,
	}
}

func ToDomainDetailList(list []Cafe) []domain.Cafe {
	results := make([]domain.Cafe, len(list))
	for i, detail := range list {
		results[i] = domain.NewCafeBuilder().
			Id(detail.Id).
			OwnerId(detail.OwnerId).
			Name(detail.Name).
			Description(detail.Description).
			CreatedAt(detail.CreatedAt).
			Build()
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
		results[i] = domain.NewCafeBuilder().
			Id(cafe.Id).
			Name(cafe.Name).
			Description(cafe.Description).
			Build()

	}
	return results
}
