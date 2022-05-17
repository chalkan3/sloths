package sloth

// import (
// 	"time"

// 	"github.com/go-kit/kit/metrics"
// )

// type instrumentingMiddleware struct {
// 	requestCount   metrics.Counter
// 	requestLatency metrics.Histogram
// 	next           Service
// }

// func NewInstrumentingMiddleware(requestCount metrics.Counter, requestLatency metrics.Histogram, next Service) *instrumentingMiddleware {
// 	return &instrumentingMiddleware{
// 		requestCount:   requestCount,
// 		requestLatency: requestLatency,
// 		next:           next,
// 	}
// }

// func (mw logmw) Add(sloth *Sloth) *Sloth {
// 	defer func(begin time.Time) {

// 	}(time.Now())
// 	return mw.svc.Add(sloth)
// }
