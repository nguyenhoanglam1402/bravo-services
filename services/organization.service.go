package services

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	payload_struct "bravo-service/api/structs"
	"errors"
)

func CreateOrganization(pl *payload_struct.SCreateOrganPayload, uid string) (string, error) {
	orgModel := model.SOrganizationModel{
		Name:    pl.Name,
		ZipCode: pl.ZipCode,
		Address: pl.Address,
		Country: pl.Country,
		Email:   pl.Email,
		OwnerId: uid,
	}

	if err := database.DB.Create(&orgModel); err != nil {
		return "", errors.New("cannot create this organization")
	}

	return "create user successfully", nil
}
