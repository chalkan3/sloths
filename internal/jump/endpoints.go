package jump

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type JumpRequest struct {
	Pos  int    `json:"pos,omitempty"`
	Name string `json:"name,omitempty"`
}

type JumpRespose struct{}

func JumpEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(JumpRequest)
		svc.Jump(request)
		return JumpRespose{}, nil
	}
}
