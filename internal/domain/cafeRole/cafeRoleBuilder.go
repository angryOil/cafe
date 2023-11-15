package cafeRole

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CafeId(cafeId int) Builder
	Name(name string) Builder
	Description(description string) Builder

	Build() CafeRole
}

type builder struct {
	id          int
	cafeId      int
	name        string
	description string
}

func (b *builder) Id(id int) Builder {
	b.id = id
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

func (b *builder) Build() CafeRole {
	return &cafeRole{
		id:          b.id,
		cafeId:      b.cafeId,
		name:        b.name,
		description: b.description,
	}
}
