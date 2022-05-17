package jump

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
			Subsystem: "jump",
			Name:      "events_jump",
			Help:      "Number of events jump received.",
		}, []string{"method", "type"}),
	}
}

func (m *Metrics) EventsFetched() *kitprometheus.Counter {
	return m.eventsFetched
}
