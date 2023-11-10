package boardType

import "time"

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CreateBy(createBy int) Builder
	CafeId(cafeId int) Builder
	Name(name string) Builder
	Description(description string) Builder
	CreatedAt(createdAt time.Time) Builder
	Build() BoardType
}

type builder struct {
	id          int
	createBy    int
	cafeId      int
	name        string
	description string
	createdAt   time.Time
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) CreateBy(createBy int) Builder {
	b.createBy = createBy
	return b
}

func (b *builder) CafeId(cafeId int) Builder {
	b.cafeId = cafeId
	return b
}

func (b *builder) Name(name string) Builder {
	b.name = name
	return b
}

func (b *builder) Description(description string) Builder {
	b.description = description
	return b
}

func (b *builder) CreatedAt(createdAt time.Time) Builder {
	b.createdAt = createdAt
	return b
}

func (b *builder) Build() BoardType {
	return &boardType{
		id:          b.id,
		createBy:    b.createBy,
		cafeId:      b.cafeId,
		name:        b.name,
		description: b.description,
		createdAt:   b.createdAt,
	}
}
