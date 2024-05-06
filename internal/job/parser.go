package job

import (
	"time"

	"github.com/gitploy-io/cronexpr"
)

var NextExecution = func(expr string) time.Time {
	nextTime := cronexpr.MustParse(expr).Next(time.Now())

	return nextTime
}
