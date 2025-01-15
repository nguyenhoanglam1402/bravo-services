package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterOrganization(c *gin.Context) {
	var pldRegister payload_struct.SCreateOrganPayload

	if err := c.ShouldBind(&pldRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload structure"})
	}

	data, err := services.CreateOrganization(&pldRegister, c.GetHeader("uid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot create organization"})
	}
	c.JSON(http.StatusOK, gin.H{data: data})
}
