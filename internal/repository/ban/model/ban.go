package model

import (
	"cafe/internal/domain/ban"
	"cafe/internal/repository/ban/req"
	"github.com/uptrace/bun"
	"time"
)

type Ban struct {
	bun.BaseModel `bun:"table:ban,alias:b"`

	Id          int       `bun:"id,pk,autoincrement"`
	UserId      int       `bun:"user_id,notnull"`
	MemberId    int       `bun:"member_id,notnull"`
	CafeId      int       `bun:"cafe_id,notnull"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at"`
}

func ToCreateModel(c req.Create) Ban {
	return Ban{
		UserId:      c.UserId,
		MemberId:    c.MemberId,
		CafeId:      c.CafeId,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
	}
}

func ToBanDomainList(list []Ban) []ban.Ban {
	results := make([]ban.Ban, len(list))
	for i, b := range list {
		results[i] = ban.NewBuilder().
			Id(b.Id).
			UserId(b.UserId).
			MemberId(b.MemberId).
			CafeId(b.CafeId).
			Description(b.Description).
			CreatedAt(b.CreatedAt).
			Build()
	}
	return results
}
