package routers

import (
	"bravo-service/api/controllers"
	"bravo-service/api/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouters(router *gin.RouterGroup) {
	userRt := router.Group("/users")
	userRt.Use(middlewares.JWTMiddleware())
	{
		userRt.GET("/")
		userRt.GET("/:id")
		userRt.GET("/self/profile", controllers.GetProfileHandler)
	}
}
