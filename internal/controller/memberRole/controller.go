package memberRole

import "cafe/internal/service/memberRole"

type Controller struct {
	s memberRole.Service
}

func NewController(s memberRole.Service) Controller {
	return Controller{s: s}
}
