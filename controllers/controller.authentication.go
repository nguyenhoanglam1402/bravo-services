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
		resp := payload_struct.SRespPayload{
			Message: "Invalid payload structure",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	res, err := services.LoginService(&bodyPld)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp := payload_struct.SRespPayload{
		Body: map[string]interface{}{"access_token": res},
	}

	c.JSON(http.StatusOK, resp)
}

func SignUpHandler(c *gin.Context) {
	var bodyPld payload_struct.SSignUpPayload

	if err := c.ShouldBind(&bodyPld); err != nil {
		resp := payload_struct.SRespPayload{
			Message: "Invalid payload structure",
			Detail:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := services.SignUpService(&bodyPld)

	if err != nil {
		resp := payload_struct.SRespPayload{
			Message: "Unexpected Error",
			Detail:  err.Error(),
		}
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	successResp := payload_struct.SRespPayload{
		Body: map[string]interface{}{
			"access_token": data,
		},
	}

	c.JSON(http.StatusOK, successResp)
}
