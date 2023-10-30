package repository

import (
	"cafe/internal/domain"
	page2 "cafe/internal/page"
	"cafe/internal/repository/model"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
)

type BanRepository struct {
	db bun.IDB
}

func NewBanRepository(db bun.IDB) BanRepository {
	return BanRepository{
		db: db,
	}
}

func (r BanRepository) Create(ctx context.Context, bDomain domain.Ban) error {
	bModel := model.ToBanModel(bDomain)
	_, err := r.db.NewInsert().Model(&bModel).Exec(ctx)
	return err
}

func (r BanRepository) GetListCountByUserId(ctx context.Context, userId int, reqPage page2.ReqPage) ([]domain.Ban, int, error) {
	var results []model.Ban
	count, err := r.db.NewSelect().Model(&results).Where("user_id = ?", userId).Limit(reqPage.Size).Offset(reqPage.Offset).Order("id desc").ScanAndCount(ctx)
	if err != nil {
		log.Println("GetListCountByUserId NewSelect err: ", err)
		return []domain.Ban{}, 0, errors.New("internal server error")
	}
	return model.ToBanDomainList(results), count, nil
}

func (r BanRepository) GetListCountByCafeId(ctx context.Context, cafeId int, reqPage page2.ReqPage) ([]domain.Ban, int, error) {
	var results []model.Ban
	count, err := r.db.NewSelect().Model(&results).Where("cafe_id = ?", cafeId).Limit(reqPage.Size).Offset(reqPage.Offset).Order("id desc").ScanAndCount(ctx)
	if err != nil {
		log.Println("GetListCountByCafeId NewSelect err: ", err)
		return []domain.Ban{}, 0, errors.New("internal server error")
	}
	return model.ToBanDomainList(results), count, nil
}
