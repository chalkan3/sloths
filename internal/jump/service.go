package jump

type Service interface {
	Jump(jumpRequest JumpEvent) error
}

type service struct {
	repository Repository
}

func (s *service) Jump(jumpRequest JumpEvent) error {
	s.repository.Insert(jumpRequest)
	return nil
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}
