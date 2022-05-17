package ballet

import (
	"time"

	"github.com/go-kit/log"
)

type loggerMW struct {
	logger log.Logger
	next   Service
}

func NewLoggerMW(logger log.Logger, next Service) *loggerMW {
	return &loggerMW{
		logger: logger,
		next:   next,
	}
}

func (mw *loggerMW) Dance(request DanceBalletRequest) error {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Dance Ballet",
			"type", "make a sloth dance ballet",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Dance(request)
}
