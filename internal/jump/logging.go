package jump

import (
	"time"

	"github.com/go-kit/kit/log"
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

func (mw *loggerMW) Jump(jumpRequest JumpEvent) error {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Jump",
			"type", "make a sloth jump",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.Jump(jumpRequest)
}
