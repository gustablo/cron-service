package config

import (
	"database/sql"
	"fmt"
	"github.com/gustablo/cron-service/context"
	_ "github.com/lib/pq"
	"log"
)

func NewDB() *sql.DB {
	env := context.GetContext().Env

	DB_USER := env.Get("DB_USER")
	DB_NAME := env.Get("DB_NAME")
	DB_PASS := env.Get("DB_PASS")
	DB_HOST := env.Get("DB_HOST")
	DB_PORT := env.Get("DB_PORT")
	DB_SSL_MODE := env.Get("DB_SSL_MODE")

	stringConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME, DB_SSL_MODE)

	db, err := sql.Open("postgres", stringConn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
