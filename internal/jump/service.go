package jump

import (
	"github.com/go-kit/log"
)

type Service interface {
	Jump(jumpRequest JumpEvent) error
}

type service struct {
}

func (s *service) Jump(jumpRequest JumpEvent) error {
	return nil
}

func NewService(logger log.Logger) Service {
	return new(service)
}
