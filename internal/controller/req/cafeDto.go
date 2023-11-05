package req

import (
	"cafe/internal/domain"
	"errors"
	"time"
)

type CreateCafeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c CreateCafeDto) ToDomain(userId int) (domain.Cafe, error) {
	err := validateCreateCafe(c.Name, userId)
	if err != nil {
		return domain.NewCafeBuilder().Build(), err
	}
	return domain.NewCafeBuilder().
		Name(c.Name).
		OwnerId(userId).
		Description(c.Description).
		CreatedAt(time.Now()).
		Build(), nil

}

type UpdateCafeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c UpdateCafeDto) ToDomain(userId, cafeId int) (domain.Cafe, error) {
	if c.Name == "" {
		return domain.NewCafeBuilder().Build(), errors.New("name is empty")
	}

	return domain.NewCafeBuilder().
		Id(cafeId).
		Name(c.Name).
		OwnerId(userId).
		Description(c.Description).
		Build(), nil
}

func validateCreateCafe(name string, ownerId int) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if ownerId == 0 {
		return errors.New("owner id is zero")
	}
	return nil
}
