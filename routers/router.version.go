package routers

import (
	"bravo-service/api/controllers"

	"github.com/gin-gonic/gin"
)

func VersionControlRouter(r *gin.RouterGroup) {
	verRouter := r.Group("/version-control")

	verRouter.POST("/checkout", controllers.CheckoutBranchHandler)
}
