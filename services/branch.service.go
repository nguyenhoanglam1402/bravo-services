package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	payload_struct "bravo-service/api/structs"
	"errors"
	"fmt"

	"github.com/google/uuid"
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
		fmt.Printf(branchResDB.Error.Error())
		tx.Rollback()
		return nil, errors.New("new branch cannot be checkout")
	}

	var verPld = model.SVersion{
		BranchID:        branchPld.ID,
		AuthorID:        pld.UserID,
		ParentVersionID: existVers.ID,
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
	if err := database.DB.Where("lesson_id = ?", id).Preload("Author").Find(&branchs).Error; err != nil {
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
		BranchID: pld.BranchID,
		RawData:  pld.RawData,
		CompData: pld.CompData,
		AuthorID: uidUUID,
	}
	if res := database.DB.Create(&version); res.Error != nil {
		return errors.New("branch cannot commit version")
	}
	return nil
}
