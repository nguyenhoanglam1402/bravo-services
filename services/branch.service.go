package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	payload_struct "bravo-service/api/structs"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CheckoutBranchService(pld *payload_struct.SBranchPayload) (*model.SVersion, error) {
	var existVers model.SVersion
	if dbRes := database.DB.Where(&model.SVersion{ID: pld.VesionID}).Preload("Branch").First(&existVers); dbRes.Error != nil {
		return nil, errors.New("current version is not exist")
	}

	tx := database.DB.Begin()

	var branchPld = model.SBranch{
		Name:      pld.BranchName,
		LessonID:  existVers.Branch.LessonID,
		CreatedBy: pld.UserID,
	}

	branchResDB := tx.Create(&branchPld)

	if branchResDB.Error != nil {
		tx.Rollback()
		return nil, errors.New("new branch cannot be checkout")
	}

	var verPld = model.SVersion{
		BranchID:        branchPld.ID,
		AuthorID:        pld.UserID,
		ParentVersionID: &existVers.ID,
		RawData:         existVers.RawData,
		CompData:        existVers.CompData,
	}

	verResDB := tx.Create(&verPld)

	if verResDB.Error != nil {
		tx.Rollback()
		return nil, errors.New("new branch cannot be checkout")
	}

	tx.Commit()

	return &verPld, nil
}

func GetListBranchService(id uuid.UUID) ([]model.SBranch, error) {
	var branchs []model.SBranch
	// Only select necessary fields and preload minimal author info
	if err := database.DB.Select("id, name, lesson_id, created_by, created_at").Where("lesson_id = ?", id).
		Preload("Author", func(db *gorm.DB) *gorm.DB { return db.Select("id, fullname, email") }).
		Find(&branchs).Error; err != nil {
		return nil, errors.New("cannot find any lesson with this id")
	}
	return branchs, nil
}

func CommitBranchService(pld *payload_struct.SCommitBranchPayload, uid string) error {
	var uidUUID uuid.UUID
	var uuidErr error

	if uidUUID, uuidErr = uuid.Parse(uid); uuidErr != nil {
		return errors.New("UID is not authorized")
	}

	version := model.SVersion{
		BranchID:        pld.BranchID,
		RawData:         pld.RawData,
		CompData:        pld.CompData,
		Message:         pld.MessageContent,
		AuthorID:        uidUUID,
		ParentVersionID: pld.ParentVersionID,
	}
	if res := database.DB.Create(&version); res.Error != nil {
		fmt.Printf("Error: %s", res.Error.Error())
		return errors.New("branch cannot commit version")
	}
	return nil
}

func GetAllVersionBranchService(branchId uuid.UUID) (*[]payload_struct.SVersionBranchTreePayload, error) {
	var versions []model.SVersion
	var versionParent []model.SVersion
	// Only select necessary fields for performance
	if err := database.DB.Select("id, branch_id, author_id, parent_version_id, raw_data, comp_data, message, created_at").Where("branch_id = ?", branchId).Find(&versions).Error; err != nil {
		return nil, errors.New("not found any record")
	}
	if len(versions) == 0 {
		return nil, errors.New("no versions found for this branch")
	}
	parentBranchId := versions[0].ParentVersionID
	if parentBranchId == nil {
		parentBranchId = &branchId // fallback to self if no parent
	}
	if err := database.DB.Select("id, branch_id, author_id, parent_version_id, raw_data, comp_data, message, created_at").Where("branch_id = ?", parentBranchId).Find(&versionParent).Error; err != nil {
		return nil, errors.New("not found any parent record")
	}
	branchIds := []uuid.UUID{branchId, *parentBranchId}
	var branchs []payload_struct.SVersionBranchTreePayload
	for _, id := range branchIds {
		var versionContents []model.SVersion
		if id == branchId {
			versionContents = versions
		} else {
			versionContents = versionParent
		}
		branchs = append(branchs, payload_struct.SVersionBranchTreePayload{
			BranchID:   id,
			BranchName: "", // Optionally load name if needed
			Versions:   &versionContents,
		})
	}
	return &branchs, nil
}
