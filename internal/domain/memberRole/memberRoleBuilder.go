package memberRole

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CafeId(cafeId int) Builder
	CafeRoleIds(cafeRoleIds string) Builder
	MemberId(memberId int) Builder
	Build() MemberRole
}

type builder struct {
	id          int
	cafeId      int
	cafeRoleIds string
	memberId    int
}

func (b *builder) CafeId(cafeId int) Builder {
	b.cafeId = cafeId
	return b
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) CafeRoleIds(cafeRoleIds string) Builder {
	b.cafeRoleIds = cafeRoleIds
	return b
}

func (b *builder) MemberId(memberId int) Builder {
	b.memberId = memberId
	return b
}

func (b *builder) Build() MemberRole {
	return &memberRole{
		id:          b.id,
		cafeRoleIds: b.cafeRoleIds,
		memberId:    b.memberId,
	}
}
