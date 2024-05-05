package job

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ExecutionTime time.Time
	LastRun       time.Time
	Uuid          string
	Expression    string
	Name          string
}

func NewJob(name string, expression string) *Job {
	nextExecution := NextExecution(expression)

	return &Job{
		Uuid:          uuid.NewString(),
		ExecutionTime: nextExecution,
		Name:          name,
		Expression:    expression,
		LastRun:       nextExecution,
	}
}

func (job *Job) IsJobScheduledBefore(job2 *Job) bool {
	return job.ExecutionTime.Before(job2.ExecutionTime)
}
