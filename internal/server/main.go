package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustablo/cron-service/internal/server/controllers"
)

type Server struct {
	Addr string
}

func NewServer() *Server {
	return &Server{
		Addr: ":8080",
	}
}

func (s *Server) ServeHTTP() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1 := r.Group("/api/v1")

	jobs := v1.Group("/jobs")
	jobs.POST("", controllers.CreateJob)
	jobs.GET("", controllers.AllJobsByUserID)

	r.Run(s.Addr)
}
