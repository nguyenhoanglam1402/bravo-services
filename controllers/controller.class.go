package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClassHandler(c *gin.Context) {
	var pld payload_struct.SCreateClassPayload

	if err := c.ShouldBind(&pld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload structure"})
		return
	}

	msg, err := services.CreateClass(&pld)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func UpdateClassHandler(c *gin.Context) {
	classID := c.Param("id")
	var pld payload_struct.SCreateClassPayload

	if err := c.ShouldBind(&pld); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload structure"})
		return
	}

	msg, err := services.UpdateClass(classID, &pld)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func DeleteClassHandler(c *gin.Context) {
	classID := c.Param("id")
	msg, err := services.DeleteClass(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
