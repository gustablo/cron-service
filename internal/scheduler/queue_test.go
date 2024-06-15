package scheduler

import (
	"testing"

	"github.com/gustablo/cron-service/internal/job"
)

type nodeInput struct {
	value *job.Job
	next  *node
}

func TestNewNode(t *testing.T) {
	job := &job.Job{}

	r := newNode(job, nil)

	if job != r.value {
		t.Errorf("got %v, want %v for value", job, r.value)
	}

	if r.next != nil {
		t.Errorf("got %v, want nil for next", r.next)
	}
}

func TestNewJobsQueue(t *testing.T) {
	r := NewJobsQueue()

	if r.head != nil {
		t.Errorf("got %v, want nil for head", r.head)
	}

	if *r.count != 0 {
		t.Errorf("got %d, want 0 for count", *r.count)
	}
}

func TestInsert(t *testing.T) {
}

func TestShift(t *testing.T) {}

func TestTail(t *testing.T) {}

func TestHead(t *testing.T) {}

func TestCount(t *testing.T) {}

func TestRemoveAt(t *testing.T) {}
