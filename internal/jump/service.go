package jump

import (
	"github.com/go-kit/log"
)

type Service interface {
	Jump(jumpRequest JumpRequest) error
}

type service struct {
	logger log.Logger
}

func (s *service) Jump(jumpRequest JumpRequest) error {
	s.logger.Log(
		jumpRequest.Name, "jumped",
		"at pos", jumpRequest.Pos)
	return nil
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}
