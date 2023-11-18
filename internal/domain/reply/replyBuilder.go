package reply

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CafeId(cafeId int) Builder
	BoardId(boardId int) Builder
	Writer(writer int) Builder
	Content(content string) Builder
	CreatedAt(createdAt string) Builder
	LastUpdatedAt(lastUpdatedAt string) Builder

	Build() Reply
}

type builder struct {
	id            int
	cafeId        int
	boardId       int
	writer        int
	content       string
	createdAt     string
	lastUpdatedAt string
}

func (b *builder) Id(id int) Builder {
	b.id = id
	return b
}

func (b *builder) CafeId(cafeId int) Builder {
	b.cafeId = cafeId
	return b
}

func (b *builder) BoardId(boardId int) Builder {
	b.boardId = boardId
	return b
}

func (b *builder) Writer(writer int) Builder {
	b.writer = writer
	return b
}

func (b *builder) Content(content string) Builder {
	b.content = content
	return b
}

func (b *builder) CreatedAt(createdAt string) Builder {
	b.createdAt = createdAt
	return b
}

func (b *builder) LastUpdatedAt(lastUpdatedAt string) Builder {
	b.lastUpdatedAt = lastUpdatedAt
	return b
}

func (b *builder) Build() Reply {
	return &reply{
		id:            b.id,
		cafeId:        b.cafeId,
		boardId:       b.boardId,
		writer:        b.writer,
		content:       b.content,
		createdAt:     b.createdAt,
		lastUpdatedAt: b.lastUpdatedAt,
	}
}
