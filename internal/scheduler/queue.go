package scheduler

import "github.com/gustablo/cron-service/internal/job"

type node struct {
	value *job.Job
	next  *node
}

func newNode(value *job.Job, next *node) *node {
	return &node{
		value: value,
		next:  next,
	}
}

type JobsQueue struct {
	head  *node
	tail  *node
	count *int
}

func NewJobsQueue() *JobsQueue {
	var count int
	return &JobsQueue{
		head:  nil,
		count: &count,
	}
}

func (l *JobsQueue) Insert(e *job.Job) {
	*l.count = *l.count + 1
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

		// if *previous* is nil it means it's the first item so the newNode is now the head
		// else the previous node points to the newNode and the newNode points to the *current*
		// *previous means it's the closest node that is supposed to be executed before the newNode*
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
}

func (l *JobsQueue) Shift() *job.Job {
	var job *job.Job

	if l.head != nil {
		job = l.head.value
		l.head = l.head.next
		if l.head == nil { // if the head is nil now it means it was the only node in the list, so tail is deleted too
			l.tail = nil
		}
		*l.count = *l.count - 1
	}

	return job
}

func (l *JobsQueue) Tail() *job.Job {
	if l.tail != nil {
		return l.tail.value
	}
	return nil
}

func (l *JobsQueue) Head() *job.Job {
	if l.head != nil {
		return l.head.value
	}
	return nil
}

func (l *JobsQueue) Count() int {
	return *l.count
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
			if current.next == nil { // if the current.next is nil now it means it was the last node in the list, so we set tails as the previous
				l.tail = previous
			}
		}

		*l.count = *l.count - 1
	}
}
