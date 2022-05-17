package ballet

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

func (mw instrumentingMiddleware) Dance(danceBalletEvent DanceBalletEvent) error {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "Dance", "type", "dance Ballet event"}
		mw.metrics.EventsFetched().With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Dance(danceBalletEvent)
}
