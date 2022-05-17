package sloth

import (
	"time"

	"github.com/go-kit/kit/log"
)

type logmw struct {
	logger log.Logger
	svc    Service
}

func NewLogMW(logger log.Logger, next Service) logmw {
	return logmw{logger, next}
}

func (mw logmw) Add(sloth *Sloth) *Sloth {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Add",
			"type", "Adding a new sloth",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Add(sloth)
}

func (mw logmw) Update(sloth *Sloth) *Sloth {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Update",
			"type", "Updating a sloth",
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.svc.Update(sloth)
}

func (mw logmw) Delete(sloth *Sloth) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Delete",
			"type", "Deleting a sloth :(",
			"took", time.Since(begin),
		)
	}(time.Now())

	mw.svc.Delete(sloth)
}

func (mw logmw) Get(sloth *Sloth) *Sloth {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Get",
			"type", "fetching a sloth",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.svc.Get(sloth)
}

func (mw logmw) List() []*Sloth {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "Get",
			"type", "fetching a sloth",
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.svc.List()
}
