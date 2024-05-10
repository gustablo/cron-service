package job

import (
	"github.com/gustablo/cron-service/context"
)

func All() ([]Job, error) {
	var jobs []Job

	rows, err := context.GetContext().DB.Query("SELECT uuid, execution_time, last_run, expression, name, webhook_url, user_id FROM jobs ORDER BY execution_time ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.Uuid, &job.ExecutionTime, &job.LastRun, &job.Expression, &job.Name, &job.WebhookURL, &job.UserID); err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func AllByUserID(userID int) ([]Job, error) {
	var jobs []Job

	rows, err := context.GetContext().DB.Query("SELECT uuid, execution_time, last_run, expression, name, webhook_url, user_id FROM jobs WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var job Job
		if err := rows.Scan(&job.Uuid, &job.ExecutionTime, &job.LastRun, &job.Expression, &job.Name, &job.WebhookURL, &job.UserID); err != nil {
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
	_, err := context.GetContext().DB.Exec("INSERT INTO jobs (uuid, execution_time, last_run, expression, name, webhook_url, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		job.Uuid, job.ExecutionTime, job.LastRun, job.Expression, job.Name, job.WebhookURL, job.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (job *Job) Update() error {
	query := "UPDATE jobs SET execution_time = $1, last_run = $2 WHERE uuid = $3"
	_, err := context.GetContext().DB.Exec(query, job.ExecutionTime, job.LastRun, job.Uuid)
	if err != nil {
		return err
	}

	return nil
}
