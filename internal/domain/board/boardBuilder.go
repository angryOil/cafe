package board

var _ Builder = (*builder)(nil)

func NewBuilder() Builder {
	return &builder{}
}

type Builder interface {
	Id(id int) Builder
	CafeId(cafeId int) Builder
	BoardType(boardType int) Builder
	Writer(writer int) Builder
	Title(title string) Builder
	Content(content string) Builder
	CreatedAt(createdAt string) Builder
	LastUpdatedAt(lastUpdatedAt string) Builder

	Build() Board
}

type builder struct {
	id            int
	cafeId        int
	boardType     int
	writer        int
	title         string
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

func (b *builder) BoardType(boardType int) Builder {
	b.boardType = boardType
	return b
}

func (b *builder) Writer(writer int) Builder {
	b.writer = writer
	return b
}

func (b *builder) Title(title string) Builder {
	b.title = title
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

func (b *builder) Build() Board {
	return &board{
		id:            b.id,
		cafeId:        b.cafeId,
		boardType:     b.boardType,
		writer:        b.writer,
		title:         b.title,
		content:       b.content,
		createdAt:     b.createdAt,
		lastUpdatedAt: b.lastUpdatedAt,
	}
}
