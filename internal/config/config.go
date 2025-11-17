package config

import (
	"os"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	return &Config{
		Port: getEnv("APP_PORT"),
	}
}

func getEnv(key string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return ""
}
