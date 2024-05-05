package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/job"
)

type JobsController struct {
	ctx *context.Context
}

func NewJobsController(ctx *context.Context) *JobsController {
	return &JobsController{
		ctx: ctx,
	}
}

type jobRequest struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

func (jc *JobsController) CreateJob(c *gin.Context) {
	var request jobRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Error"})
		return
	}

	job := job.NewJob(request.Name, request.Expression)
	job.Save()
	jc.ctx.Scheduler.InsertConcurrently(job)

	c.IndentedJSON(http.StatusCreated, gin.H{})
}
