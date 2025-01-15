package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var bodyPld payload_struct.SLoginPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload structure", "detail": err.Error()})
		return
	}

	res, err := services.LoginService(&bodyPld)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"access_token": res,
		}})
}

func SignUpHandler(c *gin.Context) {
	var bodyPld payload_struct.SSignUpPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload structure", "detail": err.Error()})
		return
	}

	data, err := services.SignUpService(&bodyPld)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}
