package ballet

type Service interface {
	Dance(request DanceBalletRequest) error
}

type service struct{}

func NewService() Service { return new(service) }

func (s service) Dance(request DanceBalletRequest) error {
	return nil
}
