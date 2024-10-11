package routers

import "github.com/gin-gonic/gin"

func UserRouters(router *gin.Engine) {
	userRt := router.Group("/users")
	{
		userRt.GET("/")
		userRt.GET("/:id")
	}
}
