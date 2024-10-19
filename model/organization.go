package model

import (
	"time"

	"gorm.io/gorm"
)

type SOrganizationModel struct {
	ID      string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name    string `gorm:"type:varchar(100);not null"`
	ZipCode string `gorm:"type:varchar(100);not null"`
	Address string `gorm:"type:varchar(255);not null"`
	Country string `gorm:"type:varchar(255);not null"`
	Email   string `gorm:"type:varchar(255);not null;unique"`

	CreatedAt time.Time      `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp(3);not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp(3)"`
	BannedAt  *time.Time     `gorm:"type:timestamp(3)"`
}

type SEduGroupModel struct {
	ID             string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string `gorm:"type:varchar(50);not null"`
	OrganizationID string `gorm:"type:uuid;not null"`

	CreatedAt time.Time      `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp(3);not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp(3);index"`
	BannedAt  *time.Time     `gorm:"type:timestamp(3)"`

	Organization SOrganizationModel `gorm:"foreignKey:OrganizationID;constraint:OnDelete:CASCADE"`
}

type SClassModel struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ClassName  string `gorm:"type:varchar(50);not null"`
	EduGroupID string `gorm:"type:uuid;not null"`

	CreatedAt time.Time      `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp(3);not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp(3);index"`

	EduGroup SEduGroupModel `gorm:"foreignKey:EduGroupID;constraint:OnDelete:CASCADE"`
}
