package config

import "os"

type mongoDbConfig struct {
	Username string
	Password string
}

type serverConfig struct {
	GRPCPort string
	HTTPPort string
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
			GRPCPort: getEnv("GRPC_PORT", "50051"),
			HTTPPort: getEnv("HTTP_PORT", "8080"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
