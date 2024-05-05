package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	get := os.Getenv
	DB_USER := get("DB_USER")
	DB_NAME := get("DB_NAME")
	DB_PASS := get("DB_PASS")
	DB_HOST := get("DB_HOST")
	DB_PORT := get("DB_PORT")
	DB_SSL_MODE := get("DB_SSL_MODE")

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
