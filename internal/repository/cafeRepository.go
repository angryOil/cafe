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

func (r CafeRepository) Save(
	ctx context.Context,
	ownerId int, cafeId int,
	validFunc func(results []domain.Cafe) (domain.Cafe, error),
	mergeFunc func(findDomain domain.Cafe) domain.Cafe,
	saveValidFun func(cafe domain.Cafe) error,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("save beginTx err: ", err)
		return errors.New("internal server error")
	}
	var list []model.Cafe
	err = tx.NewSelect().Model(&list).Where("id=? and owner_id=?", cafeId, ownerId).Scan(ctx)

	if err != nil {
		log.Println("save select err: ", err)
		return errors.New("internal server error")
	}

	validDomain, err := validFunc(model.ToDomainDetailList(list))
	if err != nil {
		return err
	}

	mergedDomain := mergeFunc(validDomain)
	err = saveValidFun(mergedDomain)

	m := model.ToModel(mergedDomain)
	if err != nil {
		log.Println("save saveValidFunc err: ", err)
	}
	_, err = tx.NewInsert().Model(&m).On("conflict (id) do update").Exec(ctx)
	if err != nil {
		return err
	}
	err = tx.Commit()
	return err
}

func (r CafeRepository) GetCafesByCafeIds(ctx context.Context, ids []int) ([]domain.Cafe, error) {
	var cModels []model.Cafe
	err := r.db.NewSelect().Model(&cModels).Where("id in (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		log.Println("GetCafesByCafeIds Scan err: ", err)
		return []domain.Cafe{}, errors.New("internal server error")
	}
	return model.ToDomainDetailList(cModels), nil
}
