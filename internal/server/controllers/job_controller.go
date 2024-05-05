package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/job"
)

type jobRequest struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

func CreateJob(c *gin.Context) {
	var request jobRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Error parsing request body"})
		return
	}

	newJob := job.NewJob(request.Name, request.Expression)
	if err := newJob.Save(); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating job"})
		return
	}

	context.GetContext().Scheduler.InsertConcurrently(newJob)

	c.IndentedJSON(http.StatusCreated, gin.H{})
}
