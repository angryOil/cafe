package boardAction

import "cafe/internal/service/boardAction"

type Controller struct {
	s boardAction.Service
}

func NewController(s boardAction.Service) Controller {
	return Controller{s: s}
}
