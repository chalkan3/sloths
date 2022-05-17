package sloth

import (
	"encoding/json"

	"chalkan.github.com/internal/ballet"
	"chalkan.github.com/internal/jump"
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
}

func NewService(repository Repository, nc *nats.Conn) Service {
	return &service{
		repository: repository,
		nc:         nc,
	}
}

func (s *service) Add(sloth *Sloth) *Sloth {
	sloth = s.repository.Add(sloth)
	var queue string
	var event interface{}

	if sloth.Name == "Maria" {
		event = jump.JumpRequest{
			Pos:  1,
			Name: sloth.Name,
		}
		queue = "sloth"
	}

	if sloth.Name == "Lady" {
		event = ballet.DanceBalletRequest{
			Name:  sloth.Name,
			Music: "classic",
		}

		queue = "sloth.dance.ballet"
	}

	marshalEvent, _ := json.Marshal(event)

	s.nc.Publish(queue, marshalEvent)
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
