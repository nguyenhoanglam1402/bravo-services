package seeder

import (
	database "bravo-service/api/internal/database/config"
	"bravo-service/api/model"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedAll() {
	database.InitDatabase()
	db := database.DB
	fmt.Println("🌱 Starting database seeding...")

	err := db.Transaction(func(tx *gorm.DB) error {
		roles := seedRoles(tx)
		auths := seedAuthenticationsByRoles(tx, roles)
		users := seedUsers(tx, roles, auths)

		org := seedOrganization(tx, users[0].ID)

		// Create an edu group for the organization
		eduGroup := model.SEduGroupModel{
			ID:             uuid.New().String(),
			Name:           "Bravo Edu Group",
			OrganizationID: org.ID,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		tx.Create(&eduGroup)

		// Seed class for the edu group
		seedClass(tx, eduGroup.ID)

		lesson := seedLesson(tx, uuid.MustParse(users[0].ID), uuid.MustParse(org.ID))
		branch := seedBranch(tx, lesson.ID, uuid.MustParse(users[0].ID))
		fmt.Printf("BranchId %s \n", branch.ID)
		seedVersion(tx, branch.ID, uuid.MustParse(users[0].ID))

		return nil
	})

	if err != nil {
		fmt.Println("❌ Seeding failed and rolled back:", err)
	} else {
		fmt.Println("✅ Seeding completed.")
	}
}

func seedOrganization(db *gorm.DB, ownerID string) model.SOrganizationModel {
	org := model.SOrganizationModel{
		ID:        uuid.New().String(),
		Name:      "Bravo Academy",
		ZipCode:   "70000",
		Address:   "123 Bravo St, HCMC",
		Country:   "Vietnam",
		Email:     "contact@bravo.vn",
		OwnerId:   ownerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	db.Create(&org)
	return org
}

func seedRoles(db *gorm.DB) []model.SRoleModel {
	roleNames := []string{"Admin", "Teacher", "Viewer"}
	var roles []model.SRoleModel

	for _, name := range roleNames {
		role := model.SRoleModel{
			ID:        uuid.New().String(),
			Name:      name,
			CreatedAt: time.Now(),
		}
		db.Create(&role)
		roles = append(roles, role)
	}
	return roles
}

func seedAuthenticationsByRoles(db *gorm.DB, roles []model.SRoleModel) []model.SAuthentModel {
	var auths []model.SAuthentModel
	for i, role := range roles {
		auth := model.SAuthentModel{
			ID:        uuid.New().String(),
			Username:  fmt.Sprintf("%s_user_%d", role.Name, i+1),
			Password:  "password123", // In production, hash the password
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		db.Create(&auth)
		auths = append(auths, auth)
	}
	return auths
}

func seedUsers(db *gorm.DB, roles []model.SRoleModel, auths []model.SAuthentModel) []model.SUserModel {
	var users []model.SUserModel
	for i, role := range roles {
		user := model.SUserModel{
			ID:        uuid.New().String(),
			Fullname:  fmt.Sprintf("User %d", i+1),
			Email:     fmt.Sprintf("user%d@bravo.dev", i+1),
			Country:   "Vietnam",
			RoleID:    role.ID,
			AuthID:    auths[i].ID,
			CreatedAt: time.Now(),
		}
		db.Create(&user)
		users = append(users, user)
	}
	return users
}

func seedLesson(db *gorm.DB, authorID, orgID uuid.UUID) model.SLessonModel {
	category := model.SCategoryModel{
		ID:        uuid.New().String(),
		Name:      "Default Category",
		CreatedAt: time.Now(),
	}
	db.Create(&category)

	lesson := model.SLessonModel{
		ID:         uuid.New(),
		Name:       "Sample Lesson",
		AuthorID:   authorID,
		CategoryID: uuid.MustParse(category.ID),
		CreatedAt:  time.Now(),
	}
	db.Create(&lesson)
	return lesson
}

func seedBranch(db *gorm.DB, lessonID, createdBy uuid.UUID) model.SBranch {
	branch := model.SBranch{
		ID:        uuid.New(),
		Name:      "Main Branch",
		LessonID:  lessonID,
		CreatedBy: createdBy,
		CreatedAt: time.Now(),
	}
	if err := db.Create(&branch).Error; err != nil {
		panic(fmt.Sprintf("Failed to create branch: %v", err))
	}
	return branch
}

func seedVersion(db *gorm.DB, branchID, authorID uuid.UUID) {
	version := model.SVersion{
		ID:              uuid.New(),
		BranchID:        branchID,
		AuthorID:        authorID,
		ParentVersionID: nil,
		RawData:         "{}",
		CompData:        "",
		Message:         "Initial commit",
		CreatedAt:       time.Now(),
	}
	if err := db.Create(&version).Error; err != nil {
		panic(fmt.Sprintf("Failed to create version: %v", err))
	}
}

func seedClass(db *gorm.DB, eduGroupID string) model.SClassModel {
	class := model.SClassModel{
		ID:            uuid.New().String(),
		ClassName:     "Bravo Class 1",
		EduGroupID:    eduGroupID,
		EnrollmentKey: "enroll-2025",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	db.Create(&class)
	return class
}
