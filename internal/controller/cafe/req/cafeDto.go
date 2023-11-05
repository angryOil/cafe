package req

import (
	"cafe/internal/domain/cafe"
	"errors"
	"time"
)

type CreateCafeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c CreateCafeDto) ToDomain(userId int) (cafe.Cafe, error) {
	err := validateCreateCafe(c.Name, userId)
	if err != nil {
		return cafe.NewCafeBuilder().Build(), err
	}
	return cafe.NewCafeBuilder().
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

func (c UpdateCafeDto) ToDomain(userId, cafeId int) (cafe.Cafe, error) {
	if c.Name == "" {
		return cafe.NewCafeBuilder().Build(), errors.New("name is empty")
	}

	return cafe.NewCafeBuilder().
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
