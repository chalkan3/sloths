package jump

import (
	"context"

	"chalkan.github.com/internal/events"
	"github.com/go-kit/kit/endpoint"
)

type JumpEvent struct {
	events.Event
	Pos  int    `json:"pos,omitempty"`
	Name string `json:"name,omitempty"`
}

func (je JumpEvent) SetPos(pos int) JumpEvent {
	je.Pos = pos
	return je
}

func (je JumpEvent) SetName(name string) JumpEvent {
	je.Name = name
	return je
}

type JumpRespose struct{}

func JumpEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(JumpEvent)
		svc.Jump(request)
		return JumpRespose{}, nil
	}
}
