// api/internal/config/config.go
package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type Config struct {
	Env      string
	Database DatabaseConfig
}

var DB *gorm.DB

func LoadConfig() Config {
	// Lấy môi trường từ ENV_SYSTEM (hoặc mặc định là "local")
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// Load .env tương ứng
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Printf("⚠️ Không tìm thấy file .env.%s, dùng fallback env hệ thống", env)
	}

	return Config{
		Env: env,
		Database: DatabaseConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			DBName:   os.Getenv("POSTGRES_DB"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}
}

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	// DB.AutoMigrate(&model.SAuthentModel{})
	// DB.AutoMigrate(&model.SUserModel{})
}
