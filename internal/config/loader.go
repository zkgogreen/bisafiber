package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// getEnvWithDefault returns environment variable value or default if not set
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvBoolWithDefault returns environment variable as bool or default if not set
func getEnvBoolWithDefault(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

func Get() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
	}
	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Driver:    getEnvWithDefault("DB_DRIVER", "mysql"),
			Host:      os.Getenv("DB_HOST"),
			Port:      os.Getenv("DB_PORT"),
			Name:      os.Getenv("DB_NAME"),
			User:      os.Getenv("DB_USER"),
			Pass:      os.Getenv("DB_PASS"),
			Tz:        getEnvWithDefault("DB_TZ", "Asia/Jakarta"),
			Charset:   getEnvWithDefault("DB_CHARSET", "utf8mb4"),
			ParseTime: getEnvBoolWithDefault("DB_PARSE_TIME", true),
			Loc:       getEnvWithDefault("DB_LOC", "Local"),
		},
	}
}
