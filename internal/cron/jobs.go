package cron

import (
	"time"

	"github.com/google/uuid"
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

type node struct {
	value *Job
	next  *node
}

func newNode(value *Job, next *node) *node {
	return &node{
		value: value,
		next:  next,
	}
}

type JobsQueue struct {
	head  *node
	tail  *node
	count int
}

func NewJobsQueue() *JobsQueue {
	return &JobsQueue{
		head:  nil,
		count: 0,
	}
}

func (l *JobsQueue) Insert(e *Job) {
	newNode := newNode(e, nil)

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		var previous *node
		current := l.head
		for current != nil {
			if e.IsJobScheduledBefore(current.value) {
				break
			} else {
				previous = current
				current = current.next
			}
		}

		// if *previous* is nil it means its the first item so the newNode is now the head
		// else the previous node points to the newNode and the newNode points to the *current*
		// *previous means its the closest node that is supposed to be executed before the newNode*
		// *current means the last found node in the list that is supposed to be executed after the newNode*
		if previous == nil {
			l.head = newNode
		} else {
			previous.next = newNode
		}
		newNode.next = current

		// if current is nil it means the for statement iterated till the end then current var is the last node
		// then the tail of the list is now the newNode
		if current == nil {
			l.tail = newNode
		}
	}

	l.count++
}

func (l *JobsQueue) Shift() *Job {
	var job *Job

	if l.head != nil {
		job = l.head.value
		l.head = l.head.next
		if l.head == nil { // if the head is nil now it means it was the only node in the list, so tail is deleted too
			l.tail = nil
		}
		l.count--
	}

	return job
}

func (l JobsQueue) Iterate() {
	current := l.head

	for current != nil {
		current = current.next
	}
}

func (l JobsQueue) Tail() *Job {
	if l.tail != nil {
		return l.tail.value
	}
	return nil
}

func (l JobsQueue) Head() *Job {
	if l.head != nil {
		return l.head.value
	}
	return nil
}

func (l JobsQueue) Count() int {
	return l.count
}

func (l *JobsQueue) RemoveAt(uuid string) {
	var previous *node
	found := false
	current := l.head

	for current != nil {
		if current.value.Uuid == uuid {
			found = true
			break
		}

		previous = current
		current = current.next
	}

	if found {
		if previous == nil {
			l.head = l.head.next
			if l.head == nil { // if the head is nil now it means it was the only node in the list, so tail is deleted too
				l.tail = nil
			}
		} else {
			previous.next = current.next
			if current.next == nil { // if the current.next is nil now it means it was the last node in the list so we set tails as the previous
				l.tail = previous
			}
		}

		l.count--
	}
}

func (job *Job) IsJobScheduledBefore(job2 *Job) bool {
	return job.ExecutionTime.Before(job2.ExecutionTime)
}
