package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(envname string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err.Error())
	}

	value := os.Getenv(envname)
	return value
}
