package ballet

type Service interface {
	Dance(request DanceBalletEvent) error
}

type service struct{}

func NewService() Service { return new(service) }

func (s service) Dance(request DanceBalletEvent) error {
	return nil
}
