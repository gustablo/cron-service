package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/gustablo/cron-service/config"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/job"
)

type Scheduler struct {
	PendingQueue *JobsQueue
	RunningQueue *JobsQueue
	updateChan   chan *job.Job
	ctx          *context.Context
}

func NewScheduler(ctx *context.Context) *Scheduler {
	return &Scheduler{
		PendingQueue: NewJobsQueue(),
		RunningQueue: NewJobsQueue(),
		updateChan:   make(chan *job.Job),
		ctx:          ctx,
	}
}

func (c *Scheduler) loadJobs() {
	jobs, err := job.All()
	if err != nil {
		log.Fatal(err)
	}

	for _, job := range jobs {
		// Make a copy of job because it will be reassigned with each loop. (golang 1.21 bug)
		tempJob := job
		c.PendingQueue.Insert(&tempJob)
	}
}

func (c *Scheduler) Start() {
	c.loadJobs()

	go c.updateJobs()

	ticker := time.NewTicker(1 * time.Second)
	for {
		if !c.isRunningQueueFull() {
			if job := c.PendingQueue.Shift(); job != nil {
				c.RunningQueue.Insert(job)
				go c.process(job)
			} else {
				<-ticker.C
			}
		} else {
			<-ticker.C
		}
	}
}

func (c *Scheduler) InsertConcurrently(newJob *job.Job) {
	longestRunningJob := c.RunningQueue.Tail()
	fmt.Println(newJob.IsJobScheduledBefore(longestRunningJob), longestRunningJob.ExecutionTime, newJob.ExecutionTime)

	if longestRunningJob != nil && newJob.IsJobScheduledBefore(longestRunningJob) {
		// it is bad cause if the number of InsertConcurrently is big there will be a lot of goroutines running
		go c.process(newJob)
	} else {
		c.PendingQueue.Insert(newJob)
	}
}

func (c *Scheduler) process(j *job.Job) {
	fmt.Println("executing", j.Name)
	if j.LastRun.Equal(j.ExecutionTime) {
		j.ExecutionTime = job.NextExecution(j.Expression)
	}

	sleepDuration := time.Until(j.ExecutionTime)
	timer := time.NewTimer(sleepDuration)
	<-timer.C

	c.execute(j)
	c.reprioritizeJob(j)
}

func (c *Scheduler) reprioritizeJob(j *job.Job) {
	// put in the pending queue again after executed
	j.LastRun = j.ExecutionTime
	j.ExecutionTime = job.NextExecution(j.Expression).Add(1 * time.Second) // adding 1 sec to prevent the job to be set to the same minute

	// we should insert the job before remove it from the running queue
	// bc it avoids longest jobs to be catch first
	c.PendingQueue.Insert(j)
	c.RunningQueue.RemoveAt(j.Uuid)

	c.updateChan <- j
}

func (c *Scheduler) execute(j *job.Job) {
	fmt.Println("finished:", j.Uuid)
}

func (c *Scheduler) updateJobs() {
	for job := range c.updateChan {
		if err := job.Update(); err != nil {
			fmt.Println("Error while updating job:", job.Uuid)
		} else {
			fmt.Println("Job updated:", job.Uuid, job.ExecutionTime)
		}
	}
}

func (c *Scheduler) isRunningQueueFull() bool {
	return *c.RunningQueue.count >= config.MAX_GO_ROUTINES
}
