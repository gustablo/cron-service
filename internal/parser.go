package cron

import (
	"github.com/gitploy-io/cronexpr"
	"time"
)

func NextExecution(expr string) time.Time {
	nextTime := cronexpr.MustParse(expr).Next(time.Now())

	return nextTime
}
