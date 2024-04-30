package main

import (
	"github.com/gustablo/cron-service/internal/cron"
)

func main() {
	c := cron.NewCron()
	c.Start()
}
