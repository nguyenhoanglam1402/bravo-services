package main

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")

	if err != nil {
		panic(err)
	}

	database.InitDatabase()
	router := gin.Default()
	rV1 := router.Group("/api/v1")
	routers.AuthenticationRouter(rV1)
	router.Run(":8080")
}
