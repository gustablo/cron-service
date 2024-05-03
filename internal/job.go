package cron

import (
	"time"

	"github.com/google/uuid"
	"github.com/gustablo/cron-service/config"
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

func All() ([]Job, error) {
	var jobs []Job

	rows, err := config.DB.Query("SELECT uuid, execution_time, last_run, expression, name FROM jobs ORDER BY execution_time ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.Uuid, &job.ExecutionTime, &job.LastRun, &job.Expression, &job.Name); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (job *Job) Save() error {
	_, err := config.DB.Exec("INSERT INTO jobs (uuid, execution_time, last_run, expression, name) VALUES ($1, $2, $3, $4, $5)",
		job.Uuid, job.ExecutionTime, job.LastRun, job.Expression, job.Name)
	if err != nil {
		return err
	}

	return nil
}

func (job *Job) Update() error {
	query := "UPDATE jobs SET execution_time = $1, last_run = $2 WHERE uuid = $3"
	_, err := config.DB.Exec(query, job.ExecutionTime, job.LastRun, job.Uuid)
	if err != nil {
		return err
	}

	return nil
}
