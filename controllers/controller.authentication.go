package controllers

import (
	payload_struct "bravo-service/api/structs/auth"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var bodyPld payload_struct.SLoginPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload structure"})
	}

	fmt.Println(bodyPld.Password)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Username: %s, Password: %s", bodyPld.Username, bodyPld.Password)})
}

func SignUpHandler(c *gin.Context) {
	var bodyPld payload_struct.SSignUpPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload structure"})
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Username: %s, Password: %s", bodyPld.Username, bodyPld.Password)})
}
