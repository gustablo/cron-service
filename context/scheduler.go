package context

import "github.com/gustablo/cron-service/internal/job"

type Scheduler interface {
	Start()
	InsertConcurrently(newJob *job.Job)
}
