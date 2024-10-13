package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	"bravo-service/api/packages/utils"
	payload_struct "bravo-service/api/structs/auth"
	"errors"

	"gorm.io/gorm"
)

func SignUpService(bodyPld *payload_struct.SSignUpPayload) (*model.SUserModel, error) {
	hashPassword, err := utils.HashPassword(bodyPld.Password)

	if err != nil {
		return nil, errors.New("cannot hashing password right now")
	}

	var enUser model.SUserModel

	tranErr := database.DB.Transaction(func(tx *gorm.DB) error {
		enAuth := model.SAuthentModel{
			Username: bodyPld.Email,
			Password: hashPassword,
		}

		if err := tx.Create(&enAuth).Error; err != nil {
			return errors.New("cannot create new account")
		}

		enUser = model.SUserModel{
			FullName: bodyPld.Fullname,
			Email:    bodyPld.Email,
			JobTitle: bodyPld.JobTitle,
			Country:  bodyPld.Country,
			AuthID:   enAuth.ID,
			Auth:     enAuth,
		}

		if err := tx.Create(&enUser).Error; err != nil {
			return errors.New("cannot setup user information")
		}

		return nil
	})

	if tranErr != nil {
		return nil, tranErr
	}

	return &enUser, nil
}
