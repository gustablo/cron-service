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
	WebhookURL    string
	UserID        int
}

func NewJob(name string, expression string, webhookURL string, userID int) *Job {
	nextExecution := NextExecution(expression)

	return &Job{
		Uuid:          uuid.NewString(),
		ExecutionTime: nextExecution,
		Name:          name,
		Expression:    expression,
		LastRun:       nextExecution,
		WebhookURL:    webhookURL,
		UserID:        userID,
	}
}

func (job *Job) IsJobScheduledBefore(job2 *Job) bool {
	return job.ExecutionTime.Before(job2.ExecutionTime)
}
