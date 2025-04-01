package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetListBranchHandler(c *gin.Context) {
	lessonIdString := c.Param("id")
	lessonId, parseErr := uuid.Parse(lessonIdString)

	if parseErr != nil {
		respErr := payload_struct.SRespPayload{
			Message: "id is not valid",
		}
		c.JSON(http.StatusBadRequest, respErr)
		return
	}

	branchs, findErr := services.GetListBranchService(lessonId)

	if findErr != nil {
		respErr := payload_struct.SRespPayload{
			Message: "error during find the lesson, please try again",
		}
		c.JSON(http.StatusInternalServerError, respErr)
		return
	}

	respSuc := payload_struct.SRespPayload{
		Body: map[string]interface{}{
			"data": branchs,
		},
	}
	c.JSON(http.StatusOK, respSuc)
}

func CheckoutBranchHandler(c *gin.Context) {
	var branchPld payload_struct.SBranchPayload
	if err := c.ShouldBind(&branchPld); err != nil {
		errResp := payload_struct.SRespPayload{
			Message: "invalid payload structur",
		}
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	verRes, err := services.CheckoutBranchService((&branchPld))

	if err != nil {
		errResp := payload_struct.SRespPayload{
			Message: "Service stop serving",
			Detail:  err.Error(),
		}
		c.JSON(http.StatusNotFound, errResp)
		return
	}

	resp := payload_struct.SRespPayload{
		Message: "Branch create successfully",
		Body: map[string]interface{}{
			"data": verRes,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func CommitOnBranchHandler(c *gin.Context) {
	var branchCommitPld payload_struct.SCommitBranchPayload

	res, exist := c.Get("uid")

	if err := c.ShouldBind(&branchCommitPld); err != nil || !exist {
		fmt.Println(err.Error())
		errResp := payload_struct.SRespPayload{
			Message: "invalid payload structur",
		}
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	if err := services.CommitBranchService(&branchCommitPld, res.(string)); err != nil {
		errResp := payload_struct.SRespPayload{
			Message: "Service stop serving",
			Detail:  err.Error(),
		}

		c.JSON(http.StatusInternalServerError, errResp)
		return
	}

	resp := payload_struct.SRespPayload{
		Message: "Commit create successfully",
	}
	c.JSON(http.StatusOK, resp)

}
