package utils

import (
	"bravo-service/api/internal/config"
	"os"
)

// GetDatabaseURL returns the complete database connection URL
func GetDatabaseURL() string {
	cfg := config.GetConfig()
	return "postgres://" + cfg.Database.User + ":" + cfg.Database.Password + "@" + cfg.Database.Host + ":" + cfg.Database.Port + "/" + cfg.Database.DBName + "?sslmode=disable"
}

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	cfg := config.GetConfig()
	return cfg.JWT.Secret
}

// GetServerPort returns the server port
func GetServerPort() string {
	cfg := config.GetConfig()
	return cfg.Server.Port
}

// GetCurrentEnvironment returns the current environment
func GetCurrentEnvironment() string {
	return config.GetEnvironment()
}

// IsProduction checks if the application is running in production mode
func IsProduction() bool {
	return config.IsProduction()
}

// IsDevelopment checks if the application is running in development mode
func IsDevelopment() bool {
	return config.IsDevelopment()
}

// IsLocal checks if the application is running in local mode
func IsLocal() bool {
	return config.IsLocal()
}

// GetEnvOrDefault gets an environment variable with a default value
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
