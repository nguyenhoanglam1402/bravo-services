package controllers

import (
	"bravo-service/api/services"
	payload_struct "bravo-service/api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LessonCreateHandler(c *gin.Context) {
	var lessonPld payload_struct.SCreateLessonPayload

	if err := c.ShouldBind(&lessonPld); err != nil {
		errRes := payload_struct.SRespPayload{
			Message: "Invalid payload",
		}
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := services.CreateLessonService(&lessonPld); err != nil {
		errRes := payload_struct.SRespPayload{
			Message: "Service stop serving",
		}
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	resp := payload_struct.SRespPayload{
		Message: "Create lesson successfully on master branch",
	}

	c.JSON(http.StatusOK, resp)
}

func LessonGetHandler(c *gin.Context) {

	branchStrId := c.Query("branch_id")

	branchId, braErr := uuid.Parse(branchStrId)

	if braErr != nil {
		errRes := payload_struct.SRespPayload{
			Message: "Invalid payload",
		}
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	lessonPld := payload_struct.SGetLessonPayload{
		BranchId: branchId,
	}

	data, err := services.GetLessonDataService(&lessonPld)

	if err != nil {
		errRes := payload_struct.SRespPayload{
			Message: "Not found the record",
		}
		c.JSON(http.StatusNotFound, errRes)
		return
	}

	resp := payload_struct.SRespPayload{
		Body: map[string]interface{}{
			"data": data,
		},
	}
	c.JSON(http.StatusOK, resp)
}
