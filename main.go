package main

import (
	"bravo-service/api/internal/config"
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/middlewares"
	"bravo-service/api/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Log current environment
	log.Printf("Starting Bravo Services in %s environment", cfg.Environment)

	// Set Gin mode
	gin.SetMode(cfg.Server.GinMode)

	// Initialize database
	database.InitDatabase()

	// Setup router
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	rV1 := router.Group("/api/v1")

	//Routers
	routers.AuthenticationRouter(rV1)
	routers.UserRouters(rV1)
	routers.VersionControlRouter(rV1)
	routers.LessionRouter(rV1)
	routers.ClassRouter(rV1)

	//Server starting
	log.Printf("Server starting on port %s in %s mode", cfg.Server.Port, cfg.Server.GinMode)
	router.Run(":" + cfg.Server.Port)
}
