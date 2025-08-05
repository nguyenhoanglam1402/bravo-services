package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Environment constants
const (
	EnvDevelopment = "dev"
	EnvLocal       = "local"
	EnvProduction  = "prod"
)

// Config holds all configuration for our application
type Config struct {
	Environment string
	Database    DatabaseConfig
	Server      ServerConfig
	JWT         JWTConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port    string
	GinMode string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret    string
	ExpiresIn string
}

var appConfig *Config

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	env := detectEnvironment()
	loadEnvironmentFile(env)

	cfg := &Config{
		Environment: env,
		Database: DatabaseConfig{
			Host:     getEnv("POSTGRES_HOST", defaultByEnv(env, "POSTGRES_HOST")),
			User:     getEnv("POSTGRES_USER", defaultByEnv(env, "POSTGRES_USER")),
			Password: getEnv("POSTGRES_PASSWORD", ""),
			DBName:   getEnv("POSTGRES_DB", defaultByEnv(env, "POSTGRES_DB")),
			Port:     getEnv("POSTGRES_PORT", defaultByEnv(env, "POSTGRES_PORT")),
		},
		Server: ServerConfig{
			Port:    getEnv("PORT", defaultByEnv(env, "PORT")),
			GinMode: getEnv("GIN_MODE", defaultByEnv(env, "GIN_MODE")),
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", defaultByEnv(env, "JWT_SECRET")),
			ExpiresIn: getEnv("JWT_EXPIRES_IN", defaultByEnv(env, "JWT_EXPIRES_IN")),
		},
	}

	appConfig = cfg
	log.Printf("Configuration loaded for environment: %s", env)
	return cfg
}

// GetConfig returns the loaded configuration
func GetConfig() *Config {
	if appConfig == nil {
		return LoadConfig()
	}
	return appConfig
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// detectEnvironment determines the current environment
func detectEnvironment() string {
	envKeys := []string{"APP_ENV", "ENVIRONMENT", "GO_ENV"}
	var env string
	for _, key := range envKeys {
		env = strings.ToLower(os.Getenv(key))
		if env != "" {
			break
		}
	}
	switch env {
	case "development", EnvDevelopment:
		return EnvDevelopment
	case "production", EnvProduction:
		return EnvProduction
	case EnvLocal:
		return EnvLocal
	default:
		log.Printf("Unknown environment '%s', defaulting to '%s'", env, EnvLocal)
		return EnvLocal
	}
}

// loadEnvironmentFile loads the appropriate .env file based on environment
func loadEnvironmentFile(env string) {
	var envFiles []string
	switch env {
	case EnvDevelopment:
		envFiles = []string{".env.dev", ".env.development", ".env"}
	case EnvProduction:
		envFiles = []string{".env.prod", ".env.production", ".env"}
	case EnvLocal:
		envFiles = []string{".env.local", ".env"}
	default:
		envFiles = []string{".env"}
	}
	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			log.Printf("Loaded environment file: %s", file)
			return
		}
	}
	log.Println("Warning: No .env file found, using system environment variables")
}

// defaultByEnv returns environment-specific default values
func defaultByEnv(env, key string) string {
	defaults := map[string]map[string]string{
		EnvLocal: {
			"POSTGRES_HOST":  "localhost",
			"POSTGRES_USER":  "postgres",
			"POSTGRES_DB":    "bravo_db_local",
			"POSTGRES_PORT":  "5432",
			"PORT":           "8080",
			"GIN_MODE":       "debug",
			"JWT_SECRET":     "local-secret-key",
			"JWT_EXPIRES_IN": "24h",
		},
		EnvDevelopment: {
			"POSTGRES_HOST":  "localhost",
			"POSTGRES_USER":  "postgres",
			"POSTGRES_DB":    "bravo_db_dev",
			"POSTGRES_PORT":  "5432",
			"PORT":           "8080",
			"GIN_MODE":       "debug",
			"JWT_SECRET":     "dev-secret-key",
			"JWT_EXPIRES_IN": "24h",
		},
		EnvProduction: {
			"POSTGRES_HOST":  "db",
			"POSTGRES_USER":  "postgres",
			"POSTGRES_DB":    "bravo_db",
			"POSTGRES_PORT":  "5432",
			"PORT":           "8080",
			"GIN_MODE":       "release",
			"JWT_SECRET":     "change-this-in-production",
			"JWT_EXPIRES_IN": "1h",
		},
	}
	if envDefaults, exists := defaults[env]; exists {
		if value, exists := envDefaults[key]; exists {
			return value
		}
	}
	if localDefaults, exists := defaults[EnvLocal]; exists {
		if value, exists := localDefaults[key]; exists {
			return value
		}
	}
	return ""
}

// Environment detection helper functions

func IsLocal() bool       { return GetConfig().Environment == EnvLocal }
func IsDevelopment() bool { return GetConfig().Environment == EnvDevelopment }
func IsProduction() bool  { return GetConfig().Environment == EnvProduction }
func GetEnvironment() string {
	return GetConfig().Environment
}
