package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	payload_struct "bravo-service/api/structs"
	"errors"
)

func CheckoutBranchService(pld *payload_struct.SBranchPayload) error {
	var existLess model.SLessonModel
	if err := database.DB.Where(&model.SLessonModel{ID: pld.LessonID}).First(&existLess); err != nil {
		return errors.New("lession is not exist")
	}
	return nil
}
