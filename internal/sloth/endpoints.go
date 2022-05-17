package sloth

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type AddRequest struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}
type AddResponse struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}

func AddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(AddRequest)
		added := svc.Add(request.Sloth)
		return AddResponse{
			Sloth: added,
		}, nil
	}
}

type UpdateRequest struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}
type UpdateResponse struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}

func UpdateEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(UpdateRequest)
		updated := svc.Update(request.Sloth)
		return UpdateResponse{
			Sloth: updated,
		}, nil
	}
}

type DeleteRequest struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}
type DeleteResponse struct {
	ID      int  `json:"id"`
	Deleted bool `json:"deleted"`
}

func DeleteEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(DeleteRequest)
		svc.Delete(request.Sloth)
		return DeleteResponse{
			ID:      request.Sloth.GetID(),
			Deleted: true,
		}, nil

	}
}

type GetRequest struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}
type GetResponse struct {
	Sloth *Sloth `json:"sloth,omitempty"`
}

func GetEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		request := rawRequest.(GetRequest)
		fetched := svc.Get(request.Sloth)
		return GetResponse{
			Sloth: fetched,
		}, nil

	}
}

type ListRequest struct {
}
type ListResponse struct {
	Sloth []*Sloth `json:"sloth,omitempty"`
}

func ListEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, rawRequest interface{}) (interface{}, error) {
		fetched := svc.List()
		return ListResponse{
			Sloth: fetched,
		}, nil

	}
}
