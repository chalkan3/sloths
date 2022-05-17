package ballet

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type DanceBalletRequest struct {
	Name  string `json:"name,omitempty"`
	Music string `json:"music,omitempty"`
}

type DanceBalletResponse struct {
}

func DanceBalletEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(DanceBalletRequest)
		if err := svc.Dance(request); err != nil {
			return nil, err
		}

		return new(DanceBalletResponse), nil
	}
}
