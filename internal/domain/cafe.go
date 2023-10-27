package domain

import (
	"errors"
	"time"
)

type Cafe struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
	CreatedAt   time.Time
}

func (c Cafe) ValidCafeFiled() error {
	if c.Name == "" {
		return errors.New("name is empty")
	}
	if c.OwnerId == 0 {
		return errors.New("owner id is zero")
	}
	if c.Id == 0 {
		return errors.New("id is zero")
	}
	return nil
}
