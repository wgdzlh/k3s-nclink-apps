package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func EnvVar(key string, defaultVal string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultVal
	}
	return value
}

func GetEnvOrExit(key string) string {
	value := EnvVar(key, "")
	if len(value) == 0 {
		log.Fatal("env: " + key + " not set.")
	}
	return value
}