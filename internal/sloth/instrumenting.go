package sloth

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

func (mw instrumentingMiddleware) Add(sloth *Sloth) *Sloth {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "add", "type", "requestsCount"}
		lvsLatency := []string{"method", "add", "type", "requestsCount"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Add(sloth)
}

func (mw instrumentingMiddleware) Update(sloth *Sloth) *Sloth {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "update", "type", "requestsCount"}
		lvsLatency := []string{"method", "update", "type", "requestsCount"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Update(sloth)
}

func (mw instrumentingMiddleware) Delete(sloth *Sloth) {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "delete", "type", "requestsCount"}
		lvsLatency := []string{"method", "delete", "type", "requestsCount"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	mw.next.Delete(sloth)
}

func (mw instrumentingMiddleware) Get(sloth *Sloth) *Sloth {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "get", "type", "requestsCount"}
		lvsLatency := []string{"method", "get", "type", "requestsCount"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.Get(sloth)
}

func (mw instrumentingMiddleware) List() []*Sloth {
	defer func(begin time.Time) {
		lvsRequests := []string{"method", "get", "type", "requestsCount"}
		lvsLatency := []string{"method", "get", "type", "requestsCount"}
		mw.metrics.latency.With(lvsLatency...).Observe(float64(time.Since(begin).Seconds()))
		mw.metrics.requestCounter.With(lvsRequests...).Add(1)

	}(time.Now())
	return mw.next.List()
}
