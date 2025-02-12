package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SLessonModel struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string         `gorm:"type:varchar(255);not null"`
	AuthorID     uuid.UUID      `gorm:"type:uuid;not null;index"`
	CategoryID   uuid.UUID      `gorm:"type:uuid;not null;index"`
	MainBranchID *uuid.UUID     `gorm:"type:uuid"`
	CreatedAt    time.Time      `gorm:"type:timestamp(3);default:current_timestamp"`
	UpdatedAt    time.Time      `gorm:"type:timestamp(3);default:current_timestamp"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	// Relationship for Eager Loading Mode
	Author     SUserModel     `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE"`
	Category   SCategoryModel `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE"`
	MainBranch *SBranch       `gorm:"foreignKey:MainBranchID;constraint:OnDelete:SET NULL"`
}
