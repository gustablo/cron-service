package main

import (
	"github.com/gustablo/cron-service/config"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/scheduler"
	"github.com/gustablo/cron-service/internal/server"
)

func main() {
	env := config.NewEnv()
	db := config.NewDB()
	srv := server.NewServer()
	sched := scheduler.NewScheduler()

	context.NewContext(sched, db, env)

	go sched.Start()
	srv.ServeHTTP()
}
