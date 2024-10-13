package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs/auth"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var bodyPld payload_struct.SLoginPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload structure", "detail": err.Error()})
	}

	fmt.Println(bodyPld.Password)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Username: %s, Password: %s", bodyPld.Username, bodyPld.Password)})
}

func SignUpHandler(c *gin.Context) {
	var bodyPld payload_struct.SSignUpPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload structure", "detail": err.Error()})
		return
	}

	data, err := services.SignUpService(&bodyPld)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	log.Printf("Data > %s", data.FullName)

	c.JSON(http.StatusOK, gin.H{"data": data})
}
