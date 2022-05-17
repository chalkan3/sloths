package sloth

import (
	"encoding/json"

	"chalkan.github.com/internal/ballet"
	"chalkan.github.com/internal/events"
	"chalkan.github.com/internal/jump"
)

type Definition struct {
	Message events.IEvent
	Queue   string
}

func (d Definition) GetQueue() string {
	return d.Queue
}

func (d Definition) GetMessageMarshalled() []byte {
	marshalEvent, _ := json.Marshal(d.Message)
	return marshalEvent
}

type Events struct {
	events map[string]Definition
}

func NewEvents() *Events {
	return &Events{
		events: map[string]Definition{
			"Maria": {
				Message: jump.JumpEvent{
					Pos:  1,
					Name: "Maria",
				},
				Queue: "sloth.jump",
			},
			"Lady": {
				Message: ballet.DanceBalletEvent{
					Music: "classic",
					Name:  "Lady",
				},
				Queue: "sloth.dance.ballet",
			},
		},
	}
}

func (e *Events) GetByName(name string) Definition {
	return e.events[name]
}
