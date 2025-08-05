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
		orgRouter.GET("/list", controllers.GetOrganizationListHandler)
		orgRouter.GET("/detail/:id", controllers.GetOrganizationDetailHandler)
		orgRouter.PUT("/update/:id", controllers.UpdateOrganizationHandler)
		orgRouter.DELETE("/delete/:id", controllers.DeleteOrganizationHandler)
		orgRouter.POST("/search", controllers.SearchOrganizationHandler)
	}
}
