package res

type ListTotalDto[T any] struct {
	Contents []T `json:"contents,omitempty"`
	Total    int `json:"total,omitempty"`
}

func NewListTotalDto[T any](contents []T, total int) ListTotalDto[T] {
	return ListTotalDto[T]{
		Contents: contents,
		Total:    total,
	}
}
