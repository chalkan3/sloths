package sloth

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(svc Service, logger log.Logger) *mux.Router {
	opt := options(logger)

	addHandler := httptransport.NewServer(
		AddEndpoint(svc),
		decodeAdd,
		encodeResponse,
		opt...,
	)

	updateHandler := httptransport.NewServer(
		UpdateEndpoint(svc),
		decodeUpdate,
		encodeResponse,
		opt...,
	)

	removeHandler := httptransport.NewServer(
		DeleteEndpoint(svc),
		decodeDelete,
		encodeResponse,
		opt...,
	)

	getHandler := httptransport.NewServer(
		GetEndpoint(svc),
		decodeGet,
		encodeResponse,
		opt...,
	)

	listHandler := httptransport.NewServer(
		ListEndpoint(svc),
		decodeList,
		encodeResponse,
		opt...,
	)

	r := mux.NewRouter()
	r.Methods("POST").Path(Add).Handler(addHandler)
	r.Methods("PUT").Path(Update).Handler(updateHandler)
	r.Methods("DELETE").Path(Delete).Handler(removeHandler)
	r.Methods("GET").Path(Get).Handler(getHandler)
	r.Methods("GET").Path(List).Handler(listHandler)
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r

}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}

func options(logger log.Logger) []httptransport.ServerOption {
	return []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}
}

func decodeAdd(ctx context.Context, r *http.Request) (interface{}, error) {
	var request AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdate(ctx context.Context, r *http.Request) (interface{}, error) {
	var request UpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	vars := mux.Vars(r)
	id := vars["id"]
	toInt, err := strconv.Atoi(id)
	if err != nil {
		return request, err
	}

	request.Sloth.SetID(toInt)

	return request, nil
}

func decodeDelete(ctx context.Context, r *http.Request) (interface{}, error) {
	var request DeleteRequest
	vars := mux.Vars(r)
	id := vars["id"]
	toInt, err := strconv.Atoi(id)
	if err != nil {
		return request, err
	}

	return DeleteRequest{Sloth: NewSloth(toInt, "", "")}, nil
}

func decodeGet(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetRequest
	vars := mux.Vars(r)
	id := vars["id"]
	toInt, err := strconv.Atoi(id)
	if err != nil {
		return request, err
	}

	return GetRequest{Sloth: NewSloth(toInt, "", "")}, nil
}

func decodeList(ctx context.Context, r *http.Request) (interface{}, error) {
	return new(ListRequest), nil
}
