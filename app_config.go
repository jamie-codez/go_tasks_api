package main

import (
	"os"
)

type AppConfig struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

var Envs = initAppConfig()

func initAppConfig() AppConfig {
	return AppConfig{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "dev_user"),
		DBPassword: getEnv("DB_PASSWORD", "Password123"),
		DBAddress:  getEnv("DB_ADDRESS", "localhost"),
		DBName:     getEnv("DB_NAME", "test_db"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
