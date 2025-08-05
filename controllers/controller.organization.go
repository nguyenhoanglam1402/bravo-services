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

func GetOrganizationListHandler(c *gin.Context) {
	orgs, err := services.GetOrganizationList()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": orgs})
}

func GetOrganizationDetailHandler(c *gin.Context) {
	orgID := c.Param("id")
	org, err := services.GetOrganizationDetail(orgID)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": org})
}

func UpdateOrganizationHandler(c *gin.Context) {
	orgID := c.Param("id")
	var pld payload_struct.SCreateOrganPayload
	if err := c.ShouldBind(&pld); err != nil {
		c.JSON(400, gin.H{"message": "invalid payload structure"})
		return
	}
	msg, err := services.UpdateOrganization(orgID, &pld)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": msg})
}

func DeleteOrganizationHandler(c *gin.Context) {
	orgID := c.Param("id")
	msg, err := services.DeleteOrganization(orgID)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": msg})
}

func SearchOrganizationHandler(c *gin.Context) {
	var req struct {
		Query string `json:"query" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message": "invalid search payload"})
		return
	}
	results, err := services.SearchOrganizations(req.Query)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": results})
}
