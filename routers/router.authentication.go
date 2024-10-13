package routers

import (
	"bravo-service/api/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRouter(router *gin.RouterGroup) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", controllers.LoginHandler)
		authRouter.POST("/sign-up", controllers.SignUpHandler)
	}
}
