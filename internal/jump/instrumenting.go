package jump

import (
	"time"
)

type instrumentingMiddleware struct {
	metrics *Metrics
	next    Service
}

func NewInstrumentingMiddleware(metrics *Metrics, next Service) *instrumentingMiddleware {
	return &instrumentingMiddleware{
		metrics: metrics,
		next:    next,
	}
}

func (mw instrumentingMiddleware) Jump(JumpEvent JumpEvent) error {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "Jump", "type", "jump event"}
		mw.metrics.EventsFetched().With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Jump(JumpEvent)
}
