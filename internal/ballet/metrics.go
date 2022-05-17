package ballet

import (
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	eventsFetched *kitprometheus.Counter
}

func NewMetrics() *Metrics {
	return &Metrics{
		eventsFetched: kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "sloth",
			Subsystem: "ballet",
			Name:      "events",
			Help:      "Number of events dance ballet received.",
		}, []string{"method", "type"}),
	}
}

func (m *Metrics) EventsFetched() *kitprometheus.Counter {
	return m.eventsFetched
}
