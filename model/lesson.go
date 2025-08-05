package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SLessonModel struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	AuthorID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"author_id"`
	CategoryID   uuid.UUID      `gorm:"type:uuid;not null;index" json:"category_id"`
	MainBranchID *uuid.UUID     `gorm:"type:uuid" json:"main_branch_id"`
	CreatedAt    time.Time      `gorm:"type:timestamp(3);default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"type:timestamp(3);default:current_timestamp" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relationship for Eager Loading Mode
	Author     SUserModel     `gorm:"foreignKey:AuthorID;constraint:OnDelete:CASCADE" json:"author"`
	Category   SCategoryModel `gorm:"foreignKey:CategoryID;constraint:OnDelete:CASCADE" json:"category"`
	MainBranch *SBranch       `gorm:"foreignKey:MainBranchID;constraint:OnDelete:SET NULL" json:"main_branch"`
}
