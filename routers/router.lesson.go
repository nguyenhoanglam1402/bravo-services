package routers

import (
	"bravo-service/api/controllers"

	"github.com/gin-gonic/gin"
)

func LessionRouter(r *gin.RouterGroup) {
	lRouter := r.Group("/lesson")
	lRouter.POST("/create", controllers.LessonCreateHandler)
	lRouter.GET("/fetch", controllers.LessonGetHandler)

}
