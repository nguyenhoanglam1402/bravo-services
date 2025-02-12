package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckoutBranchHandler(c *gin.Context) {
	var branchPld payload_struct.SBranchPayload
	if err := c.ShouldBind(&branchPld); err != nil {
		errResp := payload_struct.SRespPayload{
			Message: "invalid payload structur",
		}
		c.JSON(http.StatusBadRequest, errResp)
	}
	if err := services.CheckoutBranchService((&branchPld)); err != nil {
		errResp := payload_struct.SRespPayload{
			Message: "Service stop serving",
			Detail:  err.Error(),
		}
		c.JSON(http.StatusNotFound, errResp)
	}

}
