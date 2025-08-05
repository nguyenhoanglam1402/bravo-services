package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	payload_struct "bravo-service/api/structs"
	"fmt"

	"gorm.io/gorm"
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
		ParentVersionID: nil,
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
	// Optimize: Only select necessary fields, use index, and avoid loading large associations unless needed
	if err := database.DB.
		Select("id, branch_id, author_id, parent_version_id, raw_data, comp_data, created_at").
		Preload("Author", func(db *gorm.DB) *gorm.DB { return db.Select("id, fullname, email") }).
		Preload("Branch", func(db *gorm.DB) *gorm.DB { return db.Select("id, name, lesson_id") }).
		Where("branch_id = ?", pld.BranchId).
		Order("created_at DESC").
		First(&versionData).Error; err != nil {
		return nil, err
	}
	return &versionData, nil
}
