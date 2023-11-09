package model

import (
	"cafe/internal/domain/cafe"
	"cafe/internal/repository/cafe/request"
	"github.com/uptrace/bun"
	"time"
)

type Cafe struct {
	bun.BaseModel `bun:"table:cafe,alias:c"`

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
		Id:          uc.Id,
		OwnerId:     uc.OwnerId,
		Name:        uc.Name,
		Description: uc.Description,
		CreatedAt:   uc.CreatedAt,
	}
}

func ToDomainDetailList(list []Cafe) []cafe.Cafe {
	results := make([]cafe.Cafe, len(list))
	for i, detail := range list {
		results[i] = cafe.NewCafeBuilder().
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

func ToDomainList(list []CateList) []cafe.Cafe {
	results := make([]cafe.Cafe, len(list))
	for i, c := range list {
		results[i] = cafe.NewCafeBuilder().
			Id(c.Id).
			Name(c.Name).
			Description(c.Description).
			Build()
	}
	return results
}
