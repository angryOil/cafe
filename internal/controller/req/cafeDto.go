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

func validateCreateCafe(name string, ownerId int) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if ownerId == 0 {
		return errors.New("owner id is zero")
	}
	return nil
}
