package cafe

import "time"

var _ CafeBuilder = (*cafeBuilder)(nil)

func NewCafeBuilder() CafeBuilder {
	return &cafeBuilder{}
}

type CafeBuilder interface {
	Id(id int) CafeBuilder
	OwnerId(ownerId int) CafeBuilder
	Name(name string) CafeBuilder
	Description(description string) CafeBuilder
	CreatedAt(createdAt time.Time) CafeBuilder

	Build() Cafe
}
type cafeBuilder struct {
	id          int
	ownerId     int
	name        string
	description string
	createdAt   time.Time
}

func (c *cafeBuilder) Id(id int) CafeBuilder {
	c.id = id
	return c
}

func (c *cafeBuilder) OwnerId(ownerId int) CafeBuilder {
	c.ownerId = ownerId
	return c
}

func (c *cafeBuilder) Name(name string) CafeBuilder {
	c.name = name
	return c
}

func (c *cafeBuilder) Description(description string) CafeBuilder {
	c.description = description
	return c
}

func (c *cafeBuilder) CreatedAt(createdAt time.Time) CafeBuilder {
	c.createdAt = createdAt
	return c
}

func (c *cafeBuilder) Build() Cafe {
	return &cafe{
		id:          c.id,
		ownerId:     c.ownerId,
		name:        c.name,
		description: c.description,
		createdAt:   c.createdAt,
	}
}
