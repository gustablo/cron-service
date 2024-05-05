package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/cron-service/context"
	api "github.com/gustablo/cron-service/internal/api/controllers"
)

type Server struct {
	Addr           string
	ctx            *context.Context
	jobsController *api.JobsController
}

func NewServer(ctx *context.Context) *Server {
	return &Server{
		Addr:           ":8080",
		ctx:            ctx,
		jobsController: api.NewJobsController(ctx),
	}
}

func (s *Server) ServeHTTP() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/v1")

	jobs := v1.Group("/jobs")
	jobs.POST("", s.jobsController.CreateJob)

	r.Run(s.Addr)
}
