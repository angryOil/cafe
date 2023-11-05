package boardAction

import "cafe/internal/cli/boardAction"

type Service struct {
	r boardAction.Requester
}

func NewService(r boardAction.Requester) Service {
	return Service{r: r}
}
