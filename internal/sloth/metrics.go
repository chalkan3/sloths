package sloth

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	latency        *kitprometheus.Summary
	requestCounter *kitprometheus.Counter
}

func NewMetrics() *Metrics {
	return &Metrics{
		requestCounter: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "sloth",
			Subsystem: "sloth_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method", "type"}),

		latency: kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "sloth",
			Subsystem: "sloth_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method", "type"}),
	}
}

func (m *Metrics) GetRequestCounter() *kitprometheus.Counter {
	return m.requestCounter
}

func (m *Metrics) GetLatency() *kitprometheus.Summary {
	return m.latency
}
