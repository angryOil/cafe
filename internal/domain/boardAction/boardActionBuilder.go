package boardAction

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CafeId(cafeId int) Builder
	BoardTypeId(boardTypeId int) Builder
	ReadRoles(readRoles string) Builder
	CreateRoles(createRoles string) Builder

	Build() BoardAction
}

type builder struct {
	id          int
	cafeId      int
	boardTypeId int
	readRoles   string
	createRoles string
	updateRoles string
	updateAble  bool
	deleteRoles string
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) CafeId(cafeId int) Builder {
	b.cafeId = cafeId
	return b
}

func (b *builder) BoardTypeId(boardTypeId int) Builder {
	b.boardTypeId = boardTypeId
	return b
}

func (b *builder) ReadRoles(readRoles string) Builder {
	b.readRoles = readRoles
	return b
}

func (b *builder) CreateRoles(createRoles string) Builder {
	b.createRoles = createRoles
	return b
}

func (b *builder) Build() BoardAction {
	return &boardAction{
		id:          b.id,
		cafeId:      b.cafeId,
		boardTypeId: b.boardTypeId,
		readRoles:   b.readRoles,
		createRoles: b.createRoles,
	}
}
