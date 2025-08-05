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

func CreateClass(pl *payload_struct.SCreateClassPayload) (string, error) {
	classModel := model.SClassModel{
		ClassName:     pl.ClassName,
		EduGroupID:    pl.EduGroupID,
		EnrollmentKey: pl.EnrollmentKey,
	}

	if err := database.DB.Create(&classModel).Error; err != nil {
		return "", errors.New("cannot create this class")
	}

	return "create class successfully", nil
}

func UpdateClass(classID string, pl *payload_struct.SCreateClassPayload) (string, error) {
	var class model.SClassModel
	if err := database.DB.First(&class, "id = ?", classID).Error; err != nil {
		return "", errors.New("class not found")
	}

	class.ClassName = pl.ClassName
	class.EduGroupID = pl.EduGroupID
	class.EnrollmentKey = pl.EnrollmentKey

	if err := database.DB.Save(&class).Error; err != nil {
		return "", errors.New("cannot update this class")
	}

	return "update class successfully", nil
}

func DeleteClass(classID string) (string, error) {
	if err := database.DB.Delete(&model.SClassModel{}, "id = ?", classID).Error; err != nil {
		return "", errors.New("cannot delete this class")
	}
	return "delete class successfully", nil
}

// Optimize: Use Select to fetch only necessary fields for list/search, and add index usage hints for large tables.

func GetOrganizationList() ([]model.SOrganizationModel, error) {
	var orgs []model.SOrganizationModel
	// Only select necessary fields for list view
	if err := database.DB.Select("id, name, zip_code, address, country, email, owner_id, created_at, updated_at").Find(&orgs).Error; err != nil {
		return nil, errors.New("cannot fetch organizations")
	}
	return orgs, nil
}

func GetOrganizationDetail(orgID string) (*model.SOrganizationModel, error) {
	var org model.SOrganizationModel
	// Use index if available for id
	if err := database.DB.Where("id = ?", orgID).First(&org).Error; err != nil {
		return nil, errors.New("organization not found")
	}
	return &org, nil
}

func UpdateOrganization(orgID string, pl *payload_struct.SCreateOrganPayload) (string, error) {
	var org model.SOrganizationModel
	if err := database.DB.First(&org, "id = ?", orgID).Error; err != nil {
		return "", errors.New("organization not found")
	}
	org.Name = pl.Name
	org.ZipCode = pl.ZipCode
	org.Address = pl.Address
	org.Country = pl.Country
	org.Email = pl.Email
	if err := database.DB.Save(&org).Error; err != nil {
		return "", errors.New("cannot update this organization")
	}
	return "update organization successfully", nil
}

func DeleteOrganization(orgID string) (string, error) {
	if err := database.DB.Delete(&model.SOrganizationModel{}, "id = ?", orgID).Error; err != nil {
		return "", errors.New("cannot delete this organization")
	}
	return "delete organization successfully", nil
}

func SearchOrganizations(query string) ([]model.SOrganizationModel, error) {
	var orgs []model.SOrganizationModel
	// Use Select for performance, and ILIKE for case-insensitive search
	if err := database.DB.Select("id, name, zip_code, address, country, email, owner_id, created_at, updated_at").Where(
		"name ILIKE ? OR email ILIKE ? OR address ILIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%",
	).Find(&orgs).Error; err != nil {
		return nil, errors.New("cannot search organizations")
	}
	return orgs, nil
}
