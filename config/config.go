package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbServer   string
	DbPort     string
	DbDatabase string
	ViewPrefix string
	AppPort    string
}

func LoadConfig() Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return Config{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbServer:   os.Getenv("DB_SERVER"),
		DbPort:     os.Getenv("DB_PORT"),
		DbDatabase: os.Getenv("DB_DATABASE"),
		ViewPrefix: os.Getenv("VIEW_PREFIX"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}
