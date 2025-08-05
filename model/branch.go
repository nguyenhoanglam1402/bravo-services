package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SBranch struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(255);not null" json:"name"`
	LessonID  uuid.UUID      `gorm:"type:uuid;not null;index" json:"lesson_id"`
	CreatedBy uuid.UUID      `gorm:"type:uuid" json:"created_by"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Lesson *SLessonModel `gorm:"foreignKey:LessonID;constraint:OnDelete:CASCADE" json:"lesson"`
	Author *SUserModel   `gorm:"foreignKey:CreatedBy;contraint:OnDelete:CASCADE" json:"author_infor"`
}
