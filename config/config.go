package config

import "os"

type mongoDbConfig struct {
	Username string
	Password string
}

type serverConfig struct {
	Port string
}

type Config struct {
	MongoDb mongoDbConfig
	Server  serverConfig
}

func NewConfig() *Config {
	return &Config{
		mongoDbConfig{
			Username: getEnv("MONGO_DB_UN", ""),
			Password: getEnv("MONGO_DB_PW", ""),
		},
		serverConfig{
			Port: getEnv("PORT", "8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
