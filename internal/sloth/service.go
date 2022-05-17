package sloth

import (
	"github.com/nats-io/nats.go"
)

type Service interface {
	Add(sloth *Sloth) *Sloth
	Update(sloth *Sloth) *Sloth
	Delete(sloth *Sloth)
	Get(sloth *Sloth) *Sloth
	List() []*Sloth
}

type service struct {
	repository Repository
	nc         *nats.Conn
	events     *Events
}

func NewService(repository Repository, nc *nats.Conn, events *Events) Service {
	return &service{
		repository: repository,
		nc:         nc,
		events:     events,
	}
}

func (s *service) Add(sloth *Sloth) *Sloth {
	sloth = s.repository.Add(sloth)
	event := s.events.GetByName(sloth.Name)
	s.nc.Publish(event.GetQueue(), event.GetMessageMarshalled())
	return sloth
}

func (s *service) Delete(sloth *Sloth) {
	s.repository.Delete(sloth)
}

func (s *service) Update(sloth *Sloth) *Sloth {
	return s.repository.Update(sloth)
}

func (s *service) Get(sloth *Sloth) *Sloth {
	return s.repository.Get(sloth)
}

func (s *service) List() []*Sloth {
	return s.repository.List()
}
