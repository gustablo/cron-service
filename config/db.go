package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenConn() {
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASS")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_SSL_MODE := os.Getenv("DB_SSL_MODE")

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

	DB = db
}
