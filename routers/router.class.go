package routers

import (
	"bravo-service/api/controllers"
	"bravo-service/api/middlewares"

	"github.com/gin-gonic/gin"
)

func ClassRouter(router *gin.RouterGroup) {
	classRouter := router.Group("/class")
	classRouter.Use(middlewares.JWTMiddleware())
	{
		classRouter.POST("/create", controllers.CreateClassHandler)
		classRouter.PUT("/update/:id", controllers.UpdateClassHandler)
		classRouter.DELETE("/delete/:id", controllers.DeleteClassHandler)
	}
}
