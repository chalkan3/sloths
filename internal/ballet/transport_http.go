package ballet

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewHTTPServer() *mux.Router {

	r := mux.NewRouter()
	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	return r

}
