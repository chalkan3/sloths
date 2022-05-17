package ballet

import (
	"context"

	"chalkan.github.com/internal/events"
	"github.com/go-kit/kit/endpoint"
)

type DanceBalletEvent struct {
	events.Event
	Name  string `json:"name,omitempty"`
	Music string `json:"music,omitempty"`
}

func (dbe DanceBalletEvent) SetPos(name string) DanceBalletEvent {
	dbe.Name = name
	return dbe
}
func (dbe DanceBalletEvent) SetMusic(music string) DanceBalletEvent {
	dbe.Music = music
	return dbe
}

type DanceBalletResponse struct {
}

func DanceBalletEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(DanceBalletEvent)
		if err := svc.Dance(request); err != nil {
			return nil, err
		}

		return new(DanceBalletResponse), nil
	}
}
