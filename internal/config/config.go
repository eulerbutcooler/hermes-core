package config

import (
	"log"
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    string
	Environment string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func LoadConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/hermes"
	}
	log.Printf("Loaded Config: Port=%s", port)
	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
		LogLevel:    getEnv("LOG_LEVEL", "INFO"),
		Environment: getEnv("ENV", "development"),
	}
}
