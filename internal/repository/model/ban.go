package model

import (
	"cafe/internal/domain"
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

func ToBanModel(d domain.Ban) Ban {
	return Ban{
		Id:          d.Id,
		UserId:      d.UserId,
		MemberId:    d.MemberId,
		CafeId:      d.CafeId,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
	}
}

func ToBanDomainList(list []Ban) []domain.Ban {
	results := make([]domain.Ban, len(list))
	for i, b := range list {
		results[i] = domain.Ban{
			Id:          b.Id,
			UserId:      b.UserId,
			MemberId:    b.MemberId,
			CafeId:      b.CafeId,
			Description: b.Description,
			CreatedAt:   b.CreatedAt,
		}
	}
	return results
}
