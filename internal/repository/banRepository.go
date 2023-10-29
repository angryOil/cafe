package repository

import (
	"cafe/internal/domain"
	"cafe/internal/repository/model"
	"context"
	"github.com/uptrace/bun"
)

type BanRepository struct {
	db bun.IDB
}

func (r BanRepository) Create(ctx context.Context, bDomain domain.Ban) error {
	bModel := model.ToBanModel(bDomain)
	_, err := r.db.NewInsert().Model(&bModel).Exec(ctx)
	return err
}

func NewBanRepository(db bun.IDB) BanRepository {
	return BanRepository{
		db: db,
	}
}
