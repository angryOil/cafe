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
		return domain.Cafe{}, err
	}
	return domain.Cafe{
		Name:        c.Name,
		OwnerId:     userId,
		Description: c.Description,
		CreatedAt:   time.Now(),
	}, nil
}

type UpdateCafeDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c UpdateCafeDto) ToDomain(userId, cafeId int) (domain.Cafe, error) {
	if c.Name == "" {
		return domain.Cafe{}, errors.New("name is empty")
	}
	return domain.Cafe{
		Id:          cafeId,
		Name:        c.Name,
		OwnerId:     userId,
		Description: c.Description,
	}, nil
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
