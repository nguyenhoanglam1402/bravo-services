package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	"bravo-service/api/packages/helper"
	"bravo-service/api/packages/utils"
	payload_struct "bravo-service/api/structs"
	"errors"
	"log"

	"gorm.io/gorm"
)

func LoginService(bodyPld *payload_struct.SLoginPayload) (string, error) {
	var enAuth model.SAuthentModel
	var enUser model.SUserModel

	authErr := database.DB.Where(&model.SAuthentModel{
		Username: bodyPld.Username,
	}).Find(&enAuth).Error

	isVerified := utils.VerifyPassword(bodyPld.Password, []byte(enAuth.Password))

	if authErr != nil || !isVerified {
		return "", errors.New("cannot found the account")
	}

	userErr := database.DB.Where(&model.SUserModel{
		AuthID: enAuth.ID,
	}).Find(&enUser).Error

	log.Println(userErr)
	log.Println(enUser)

	if userErr != nil {
		return "", errors.New("user is maybe banned")
	}

	token, errToken := helper.GenerateJWTKey(enAuth.Username, enUser.Email, enUser.RoleID, enUser.ID)

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
	var enRole model.SRoleModel

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

		roleErr := database.DB.Where(&model.SRoleModel{
			ID: bodyPld.RoleId,
		}).Find(&enRole)

		if roleErr.Error != nil {
			return errors.New("role is not found")
		}

		enUser = model.SUserModel{
			Fullname: bodyPld.Fullname,
			Email:    bodyPld.Email,
			JobTitle: bodyPld.JobTitle,
			Country:  bodyPld.Country,
			AuthID:   enAuth.ID,
			RoleID:   bodyPld.RoleId,
			Auth:     enAuth,
			Role:     enRole,
		}

		if err := tx.Create(&enUser).Error; err != nil {
			return errors.New("cannot setup user information")
		}

		token, tokenErr = helper.GenerateJWTKey(enUser.Auth.Username, enUser.Email, enUser.RoleID, enUser.ID)

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
