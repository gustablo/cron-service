package main

import (
	"github.com/gustablo/cron-service/config"
	cron "github.com/gustablo/cron-service/internal"
)

func main() {
	config.LoadEnv()
	pingDB()

	runner := cron.NewScheduler()
	runner.Start()
}

func pingDB() {
	conn, err := config.OpenConn()
	if err != nil {
		panic(err)
	}
	conn.Close()
}
