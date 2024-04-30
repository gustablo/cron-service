package cron

import (
	"fmt"
	"time"

	"github.com/gustablo/cron-service/config"
)

type Cron struct {
	PendingQueue *JobsQueue
	RunningQueue *JobsQueue
	processChan  chan *Job
}

func NewCron() *Cron {
	return &Cron{
		PendingQueue: NewJobsQueue(),
		RunningQueue: NewJobsQueue(),
		processChan:  make(chan *Job),
	}
}

func (c *Cron) Start() {
	go c.PickJobs()

	for job := range c.processChan {
		go c.process(job)
	}
}

func (c *Cron) PickJobs() {
	ticker := time.NewTicker(1 * time.Second)
	for {
		if !c.isRunningQueueFull() {
			if job := c.PendingQueue.Shift(); job != nil {
				c.processChan <- job
			} else {
				<-ticker.C
			}
		} else {
			<-ticker.C
		}
	}
}

func (c *Cron) InsertConcurrently(newJob *Job) {
	longestRunningJob := c.RunningQueue.Tail()

	if longestRunningJob != nil && newJob.IsJobScheduledBefore(longestRunningJob) {
		// it is bad cause if the number of InsertConcurrently is big there will be a lot of goroutines running
		go c.process(newJob)
	} else {
		c.PendingQueue.Insert(newJob)
	}
}

func (c *Cron) isRunningQueueFull() bool {
	return c.RunningQueue.count >= config.MAX_GO_ROUTINES
}

func (c *Cron) process(job *Job) {
	fmt.Println("executing:", job.Name)
	c.RunningQueue.Insert(job)

	sleepDuration := time.Until(job.ExecutionTime)
	timer := time.NewTimer(sleepDuration)
	<-timer.C

	c.execute(job)
	c.reprioritizeJob(job)
}

func (c *Cron) reprioritizeJob(job *Job) {
	// put in the pending queue again after executed
	job.ExecutionTime = NextExecution(job.Expression).Add(1 * time.Second) // adding 1 sec to prevent the job to be set to the same minute

	// we should insert the job before remove it from the running queue
	// bc it avoids longest jobs to be catch first
	c.PendingQueue.Insert(job)
	c.RunningQueue.RemoveAt(job.Uuid)
}

func (c *Cron) execute(job *Job) {
	fmt.Println("finished:", job.Name)
}
