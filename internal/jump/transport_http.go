package jump

import (
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer(svc Service, logger log.Logger) *mux.Router {

	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r

}
