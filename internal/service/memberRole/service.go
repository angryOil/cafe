package memberRole

import "cafe/internal/cli/memberRole"

type Service struct {
	r memberRole.Requester
}

func NewService(r memberRole.Requester) Service {
	return Service{r: r}
}
