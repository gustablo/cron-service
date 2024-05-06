package job

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	original := NextExecution
	defer func() { NextExecution = original }()

	NextExecution = func(expression string) time.Time {
		return time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	}

	name := "test_case"
	expression := "* * * * *"
	job := NewJob(name, expression)

	assert.NotNil(t, job)
	assert.Equal(t, expression, job.Expression)
	assert.Equal(t, name, job.Name)

	assert.Nil(t, uuid.Validate(job.Uuid))
	assert.Equal(t, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), job.ExecutionTime)
	assert.Equal(t, job.ExecutionTime, job.LastRun)
}

func TestIsJobScheduledBefore(t *testing.T) {
	uuid := "uuid"
	expression := "* * * * *"
	name := "test_case"
	execution := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)

	var tests = []struct {
		name  string
		input *Job
		want  bool
	}{
		{"should return true", &Job{ExecutionTime: execution.Add(1 * time.Second), Name: name, Expression: expression, LastRun: execution, Uuid: uuid}, true},
		{"should return false", &Job{ExecutionTime: execution.Add(-time.Second * 1), Name: name, Expression: expression, LastRun: execution, Uuid: uuid}, false},
	}

	job := Job{
		ExecutionTime: execution,
		LastRun:       execution,
		Uuid:          uuid,
		Expression:    expression,
		Name:          name,
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := job.IsJobScheduledBefore(tt.input)
			if r != tt.want {
				t.Errorf("got %t, want %t", r, tt.want)
			}
		})
	}
}
