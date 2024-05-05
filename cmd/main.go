package main

import (
	"github.com/gustablo/cron-service/config"
	"github.com/gustablo/cron-service/context"
	"github.com/gustablo/cron-service/internal/api"
	"github.com/gustablo/cron-service/internal/cron"
)

func main() {
	config.LoadEnv()
	config.OpenConn()

	ctx := context.CreateContext()

	ctx.Register(api.NewServer(ctx), cron.NewScheduler(ctx))

	go ctx.Scheduler.Start()
	ctx.Server.ServeHTTP()
}
