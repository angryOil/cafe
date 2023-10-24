package repository

import (
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"cafe/internal/repository/model"
	"context"
	"github.com/uptrace/bun"
)

type CafeRepository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) CafeRepository {
	return CafeRepository{db: db}
}

func (r CafeRepository) Create(ctx context.Context, cd domain.Cafe) error {
	cModel := model.ToModel(cd)
	_, err := r.db.NewInsert().Model(&cModel).Exec(ctx)
	return err
}

func (r CafeRepository) GetCafes(ctx context.Context, reqPage page2.ReqPage) ([]domain.Cafe, int, error) {
	var list []model.CateList
	err := r.db.NewSelect().Model(&list).Limit(reqPage.Size).Offset(reqPage.Offset).Order("id desc").Scan(ctx)
	if err != nil {
		return []domain.Cafe{}, 0, err
	}
	count, err := r.db.NewSelect().Model(&list).Count(ctx)
	return model.ToDomainList(list), count, err
}

func (r CafeRepository) GetDetail(ctx context.Context, id int) ([]domain.Cafe, error) {
	var list []model.Cafe
	err := r.db.NewSelect().Model(&list).Where("id=?", id).Scan(ctx)
	if err != nil {
		return []domain.Cafe{}, err
	}
	return model.ToDomainDetailList(list), nil
}
