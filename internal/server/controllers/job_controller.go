package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/job"
)

type jobRequest struct {
	Name       string `json:"Name"`
	Expression string `json:"Expression"`
	WebhookURL string `json:"WebhookUrl"`
	UserID     int    `json:"UserID"`
}

type jobResponse struct {
	Uuid       string    `json:"Uuid"`
	Name       string    `json:"Name"`
	Expression string    `json:"Expression"`
	WebhookURL string    `json:"WebhookUrl"`
	UserID     int       `json:"UserID"`
	LastRun    time.Time `json:"LastRun"`
	NextRun    time.Time `json:"NextRun"`
}

func CreateJob(c *gin.Context) {
	var request jobRequest

	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Error parsing request body"})
		return
	}

	newJob := job.NewJob(request.Name, request.Expression, request.WebhookURL, request.UserID)
	if err := newJob.Save(); err != nil {
		fmt.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error creating job"})
		return
	}

	context.GetContext().Scheduler.InsertConcurrently(newJob)

	c.IndentedJSON(http.StatusCreated, gin.H{})
}

func AllJobsByUserID(c *gin.Context) {
	userID, exists := c.GetQuery("user_id")
	if !exists {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "User ID is required"})
		return
	}

	intUserID, err := strconv.Atoi(userID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid User ID format"})
		return
	}

	jobs, err := job.AllByUserID(intUserID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error listing user jobs"})
		return
	}

	var response []jobResponse
	for _, job := range jobs {
		response = append(response, jobResponse{
			Uuid:       job.Uuid,
			Name:       job.Name,
			Expression: job.Expression,
			WebhookURL: job.WebhookURL,
			UserID:     job.UserID,
			LastRun:    job.LastRun,
			NextRun:    job.ExecutionTime,
		})
	}

	c.IndentedJSON(http.StatusOK, response)
}
