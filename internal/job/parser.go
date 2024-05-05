package job

import (
	"time"

	"github.com/gitploy-io/cronexpr"
)

func NextExecution(expr string) time.Time {
	nextTime := cronexpr.MustParse(expr).Next(time.Now())

	return nextTime
}
