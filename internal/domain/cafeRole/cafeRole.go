package cafeRole

import (
	"cafe/internal/domain/cafeRole/vo"
	"errors"
)

var _ CafeRole = (*cafeRole)(nil)

type CafeRole interface {
	ValidCreate() error
	ValidUpdate() error

	ToInfo() vo.Info
}

type cafeRole struct {
	id          int
	cafeId      int
	name        string
	description string
}

const (
	InvalidCafeId = "invalid cafe id"
	InvalidName   = "invalid name"
	InvalidId     = "invalid cafeRole id"
)

func (c *cafeRole) ValidCreate() error {
	if c.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if c.name == "" {
		return errors.New(InvalidName)
	}
	return nil
}

func (c *cafeRole) ValidUpdate() error {
	if c.id < 1 {
		return errors.New(InvalidId)
	}
	if c.name == "" {
		return errors.New(InvalidName)
	}
	if c.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	return nil
}

func (c *cafeRole) ToInfo() vo.Info {
	return vo.Info{
		Id:          c.id,
		CafeId:      c.cafeId,
		Name:        c.name,
		Description: c.description,
	}
}
