package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	"bravo-service/api/packages/helper"
	"bravo-service/api/packages/utils"
	payload_struct "bravo-service/api/structs/auth"
	"errors"
	"log"

	"gorm.io/gorm"
)

func LoginService(bodyPld *payload_struct.SLoginPayload) (string, error) {
	var enAuth model.SAuthentModel

	err := database.DB.Where(&model.SAuthentModel{
		Username: bodyPld.Username,
	}).Find(&enAuth).Error

	isVerified := utils.VerifyPassword(bodyPld.Password, []byte(enAuth.Password))

	if err != nil || !isVerified {
		return "", errors.New("cannot found the account")
	}

	token, errToken := helper.GenerateJWTKey(enAuth.Username, enAuth.Username)

	if errToken != nil {
		return "", errors.New("cannot generate token")
	}

	return token, nil
}

func SignUpService(bodyPld *payload_struct.SSignUpPayload) (string, error) {
	hashPassword, err := utils.HashPassword(bodyPld.Password)

	if err != nil {
		return "", errors.New("cannot hashing password right now")
	}

	var enUser model.SUserModel
	var token string
	var tokenErr error

	tranErr := database.DB.Transaction(func(tx *gorm.DB) error {
		enAuth := model.SAuthentModel{
			Username: bodyPld.Email,
			Password: hashPassword,
		}

		if err := tx.Create(&enAuth).Error; err != nil {
			return errors.New("cannot create new account")
		}

		enUser = model.SUserModel{
			Fullname: bodyPld.Fullname,
			Email:    bodyPld.Email,
			JobTitle: bodyPld.JobTitle,
			Country:  bodyPld.Country,
			AuthID:   enAuth.ID,
			Auth:     enAuth,
		}

		if err := tx.Create(&enUser).Error; err != nil {
			return errors.New("cannot setup user information")
		}

		token, tokenErr = helper.GenerateJWTKey(enUser.Auth.Username, enUser.Email)

		if tokenErr != nil {
			log.Println(tokenErr)
			return errors.New("cannot generate token")
		}

		return nil
	})

	if tranErr != nil {
		return "", tranErr
	}

	return token, nil
}
