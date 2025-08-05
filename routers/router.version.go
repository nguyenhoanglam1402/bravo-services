package routers

import (
	"bravo-service/api/controllers"
	"bravo-service/api/middlewares"

	"github.com/gin-gonic/gin"
)

func VersionControlRouter(r *gin.RouterGroup) {
	verRouter := r.Group("/version-control")
	verRouter.Use(middlewares.JWTMiddleware())
	{
		verRouter.POST("/checkout", controllers.CheckoutBranchHandler)
		verRouter.POST("/commit", controllers.CommitOnBranchHandler)
		verRouter.GET("/log-oneline/:id", controllers.GetBranchVersions)
	}
}
