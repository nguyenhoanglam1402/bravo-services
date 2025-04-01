package routers

import (
	"bravo-service/api/controllers"
	"bravo-service/api/middlewares"

	"github.com/gin-gonic/gin"
)

func LessionRouter(r *gin.RouterGroup) {
	lRouter := r.Group("/lesson")
	lRouter.Use(middlewares.JWTMiddleware())
	{
		lRouter.POST("/create", controllers.LessonCreateHandler)
		lRouter.GET("/fetch", controllers.LessonGetHandler)
		lRouter.GET("/branchs/:id", controllers.GetListBranchHandler)
	}

}
