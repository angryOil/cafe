package domain

type IdsTotalDomain struct {
	Ids   []int
	Total int
}

func NewIdsTotalDomain(ids []int, total int) IdsTotalDomain {
	return IdsTotalDomain{
		Ids:   ids,
		Total: total,
	}
}
