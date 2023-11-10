package cafe

import (
	"cafe/internal/domain/cafe"
	"cafe/internal/domain/cafe/vo"
	page2 "cafe/internal/page"
	"cafe/internal/repository/cafe/model"
	"cafe/internal/repository/cafe/request"
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

const (
	InternalServerError = "internal server error"
)

func (r CafeRepository) Create(ctx context.Context, cd request.CreateCafe) error {
	cModel := model.ToCreateModel(cd)
	_, err := r.db.NewInsert().Model(&cModel).Exec(ctx)
	return err
}

func (r CafeRepository) GetCafes(ctx context.Context, reqPage page2.ReqPage) ([]cafe.Cafe, int, error) {
	var list []model.CateList
	count, err := r.db.NewSelect().Model(&list).Limit(reqPage.Size).Offset(reqPage.Offset).ScanAndCount(ctx)
	if err != nil {
		log.Println("GetCafes ScanAndCount err: ", err)
		return []cafe.Cafe{}, 0, errors.New("internal server error")
	}
	return model.ToDomainList(list), count, nil
}

func (r CafeRepository) GetDetail(ctx context.Context, id int) ([]cafe.Cafe, error) {
	var list []model.Cafe
	err := r.db.NewSelect().Model(&list).Where("id=?", id).Scan(ctx)
	if err != nil {
		return []cafe.Cafe{}, err
	}
	return model.ToDomainDetailList(list), nil
}

func (r CafeRepository) Save(
	ctx context.Context,
	ownerId int, cafeId int,
	validFunc func(results []cafe.Cafe) (cafe.Cafe, error),
	mergeFunc func(findDomain cafe.Cafe) (vo.UpdateCafe, error),
	saveValidFun func(cafe cafe.Cafe) error,
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

	info, err := mergeFunc(validDomain)
	if err != nil {
		return err
	}

	m := model.ToSaveModel(request.UpdateCafe{
		Id:          info.Id,
		OwnerId:     info.OwnerId,
		Name:        info.Name,
		Description: info.Description,
		CreatedAt:   info.CreatedAt,
	})
	err = saveValidFun(validDomain)
	if err != nil {
		log.Println("save saveValidFunc err: ", err)
		return err
	}
	_, err = tx.NewInsert().Model(&m).
		//Column("name").
		//Column("description").
		On("conflict (id) do update").Exec(ctx)
	if err != nil {
		log.Println("Save NewInsert err: ", err)
		return errors.New(InternalServerError)
	}
	err = tx.Commit()
	return err
}

func (r CafeRepository) GetCafesByCafeIds(ctx context.Context, ids []int) ([]cafe.Cafe, error) {
	var cModels []model.Cafe
	err := r.db.NewSelect().Model(&cModels).
		Column("id").
		Column("name").
		Where("id in (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		log.Println("GetCafesByCafeIds Scan err: ", err)
		return []cafe.Cafe{}, errors.New("internal server error")
	}
	return model.ToDomainDetailList(cModels), nil
}

func (r CafeRepository) IsExistsByUserIdCafeId(ctx context.Context, userId int, cafeId int) (bool, error) {
	ok, err := r.db.NewSelect().Model((*model.Cafe)(nil)).Where("owner_id = ? and id = ?", userId, cafeId).Exists(ctx)
	if err != nil {
		log.Println("IsExistsByUserIdCafeId err: ", err)
		return false, errors.New("internal server error")
	}
	return ok, nil
}

func (r CafeRepository) IsExistsByCafeId(ctx context.Context, cafeId int) (bool, error) {
	ok, err := r.db.NewSelect().Model((*model.Cafe)(nil)).Where("id = ?", cafeId).Exists(ctx)
	if err != nil {
		log.Println("IsExistsByCafeId select err: ", err)
		return false, errors.New("internal server error")
	}
	return ok, nil
}

func (r CafeRepository) GetOwnerIds(ctx context.Context, id int) ([]cafe.Cafe, error) {
	var m []model.Cafe
	err := r.db.NewSelect().Model(&m).
		Column("owner_id").
		Where("id = ?", id).Scan(ctx)
	if err != nil {
		log.Println("GetOwnerIds select err: ", err)
		return nil, err
	}

	return converterOwnerIds(m), nil
}

func converterOwnerIds(ms []model.Cafe) []cafe.Cafe {
	result := make([]cafe.Cafe, len(ms))
	for i, m := range ms {
		result[i] = cafe.NewCafeBuilder().OwnerId(m.OwnerId).Build()
	}
	return result
}