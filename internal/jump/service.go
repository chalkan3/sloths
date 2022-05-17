package jump

import (
	"github.com/go-kit/log"
)

type Service interface {
	Jump(jumpRequest JumpRequest) error
}

type service struct {
}

func (s *service) Jump(jumpRequest JumpRequest) error {
	return nil
}

func NewService(logger log.Logger) Service {
	return new(service)
}
