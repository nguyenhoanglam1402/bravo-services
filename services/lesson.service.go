package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	payload_struct "bravo-service/api/structs"
	"fmt"

	"github.com/google/uuid"
)

func CreateLessonService(pld *payload_struct.SCreateLessonPayload) error {
	lessonData := model.SLessonModel{
		Name:         pld.Name,
		AuthorID:     pld.AuthorId,
		CategoryID:   pld.CategoryID,
		MainBranchID: nil,
	}

	tx := database.DB.Begin()

	lessonRes := tx.Create(&lessonData)
	if lessonRes.Error != nil {
		return lessonRes.Error
	}

	branchData := model.SBranch{
		Name:      "master",
		LessonID:  lessonData.ID,
		Lesson:    &lessonData,
		CreatedBy: pld.AuthorId,
	}

	branchRes := tx.Create(&branchData)

	if branchRes.Error != nil {
		fmt.Print(branchRes.Error)
		tx.Rollback()
		return branchRes.Error
	}

	if err := tx.Model(&lessonData).Update("main_branch_id", branchData.ID).Error; err != nil {
		tx.Rollback()
		return err
	}

	versionData := model.SVersion{
		BranchID:        branchData.ID,
		AuthorID:        pld.AuthorId,
		ParentVersionID: uuid.UUID{},
		RawData:         "",
		CompData:        "",
	}

	if err := tx.Create(&versionData).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func GetLessonDataService(pld *payload_struct.SGetLessonPayload) (*model.SVersion, error) {

	var versionData model.SVersion
	if err := database.DB.Preload("Author").Preload("Branch").Where(model.SVersion{BranchID: pld.BranchId}).Order("created_at DESC").First(&versionData).Error; err != nil {
		return nil, err
	}

	return &versionData, nil
}
