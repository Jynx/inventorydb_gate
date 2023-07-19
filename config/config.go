package config

import (
	"fmt"
	"log"
	"os"
)

type mongoDbConfig struct {
	Username string
	Password string
	Protocol string
	Host     string
	Port     string
	DBName   string
}

type serverConfig struct {
	GRPCPort string
	HTTPPort string
}

type Config struct {
	MongoDb mongoDbConfig
	Server  serverConfig
}

func LoadConfig() *Config {
	mongoDbUsername, err := getEnv("MONGO_DB_UN")
	if err != nil {
		log.Fatal(err)
	}

	mongoDbPassword, err := getEnv("MONGO_DB_PW")
	if err != nil {
		log.Fatal(err)
	}

	dbName, err := getEnv("MONGO_DB_NAME")
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		mongoDbConfig{
			Username: mongoDbUsername,
			Password: mongoDbPassword,
			Protocol: getEnvWithFallback("MONGO_DB_PROTOCOL", "mongodb"),
			Host:     getEnvWithFallback("MONGO_DB_HOST", "localhost"),
			Port:     getEnvWithFallback("MONGO_DB_PORT", "27017"),
			DBName:   dbName,
		},
		serverConfig{
			GRPCPort: getEnvWithFallback("GRPC_PORT", "50051"),
			HTTPPort: getEnvWithFallback("HTTP_PORT", "8080"),
		},
	}
}

func getEnvWithFallback(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

func getEnv(key string) (string, error) {
	if value, exists := os.LookupEnv(key); exists {
		return value, nil
	}
	return "", fmt.Errorf("required env variable: %s is missing", key)
}
