package routers

import (
	"bravo-service/api/controllers"
	"bravo-service/api/middlewares"

	"github.com/gin-gonic/gin"
)

func OrganizationRouter(router *gin.RouterGroup) {
	orgRouter := router.Group("/organization")
	orgRouter.Use(middlewares.JWTMiddleware())
	{
		orgRouter.POST("/register", controllers.RegisterOrganization)
		orgRouter.PUT("/update", controllers.SignUpHandler)
	}
}
