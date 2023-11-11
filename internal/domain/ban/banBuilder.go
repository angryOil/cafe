package ban

import "time"

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	UserId(userId int) Builder
	MemberId(memberId int) Builder
	CafeId(cafeId int) Builder
	Description(description string) Builder
	CreatedAt(createdAt time.Time) Builder
	Build() Ban
}

type builder struct {
	id          int
	userId      int
	memberId    int
	cafeId      int
	description string
	createdAt   time.Time
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) UserId(userId int) Builder {
	b.userId = userId
	return b
}

func (b *builder) MemberId(memberId int) Builder {
	b.memberId = memberId
	return b
}

func (b *builder) CafeId(cafeId int) Builder {
	b.cafeId = cafeId
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

func (b *builder) Build() Ban {
	return &ban{
		id:          b.id,
		userId:      b.userId,
		memberId:    b.memberId,
		cafeId:      b.cafeId,
		description: b.description,
		createdAt:   b.createdAt,
	}
}
