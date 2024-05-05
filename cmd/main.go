package main

import (
	"github.com/gustablo/cron-service/config"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/api"
	"github.com/gustablo/cron-service/internal/cron"
)

func main() {
	env := config.NewEnv()
	db := config.NewDB()
	srv := api.NewServer()
	scheduler := cron.NewScheduler()

	context.NewContext(scheduler, db, env)

	go scheduler.Start()
	srv.ServeHTTP()
}
