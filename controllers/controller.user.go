package controllers

import (
	"bravo-service/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfileHandler(c *gin.Context) {
	email, exist := c.Get("username")

	if !exist {
		c.JSON(http.StatusForbidden, gin.H{"message": "request header authorization invalid"})
	}

	res, err := services.GetProfile(email.(string))

	if err != nil {
		c.JSON(http.StatusFound, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}
