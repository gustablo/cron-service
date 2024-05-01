package cron

import (
	"time"

	"github.com/google/uuid"
	"github.com/gustablo/cron-service/config"
)

type Job struct {
	Uuid          string
	ExecutionTime time.Time
	Expression    string
	Name          string
}

func NewJob(name string, expression string) *Job {
	return &Job{
		Uuid:          uuid.NewString(),
		ExecutionTime: NextExecution(expression),
		Name:          name,
		Expression:    expression,
	}
}

func (job *Job) IsJobScheduledBefore(job2 *Job) bool {
	return job.ExecutionTime.Before(job2.ExecutionTime)
}

func (job *Job) Save() error {
	db, err := config.OpenConn()
	if err != nil {
		return err
	}

  err = db.QueryRow("INSERT INTO jobs (id,)", args ...any)
}
