package main

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	router := gin.Default()
	rV1 := router.Group("/api/v1")
	routers.AuthenticationRouter(rV1)
	router.Run(":8080")
}
