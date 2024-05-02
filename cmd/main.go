package main

import (
	"github.com/gustablo/cron-service/config"
	cron "github.com/gustablo/cron-service/internal"
)

func main() {
	config.LoadEnv()
	config.OpenConn()
	config.Migrate()

	runner := cron.NewScheduler()
	runner.Start()
}
