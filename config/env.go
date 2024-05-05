package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct{}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{}
}

func (e Env) Get(path string) string {
	return os.Getenv(path)
}
