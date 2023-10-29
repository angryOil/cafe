package req

import (
	"cafe/internal/domain"
	"time"
)

type CreateBanDto struct {
	MemberId    int    `json:"member_id"`
	Description string `json:"description"`
}

func (d CreateBanDto) ToDomain(userId, cafeId int) domain.Ban {
	return domain.Ban{
		UserId:      userId,
		MemberId:    d.MemberId,
		CafeId:      cafeId,
		Description: d.Description,
		CreatedAt:   time.Now(),
	}
}
