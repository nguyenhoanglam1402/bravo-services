package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	"errors"

	"gorm.io/gorm/clause"
)

func GetProfile(email string) (*model.SUserModel, error) {
	var userRes model.SUserModel
	if err := database.DB.Where("email = ?", email).Preload(clause.Associations).Find(&userRes).Error; err != nil {
		return nil, errors.New("no data about this user")
	}

	return &userRes, nil
}
