package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SBranch struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null"`
	LessonID  uuid.UUID      `gorm:"type:uuid;not null;index"`
	CreatedBy uuid.UUID      `gorm:"type:uuid"`
	CreatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp;default:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Lesson *SLessonModel `gorm:"foreignKey:LessonID;constraint:OnDelete:CASCADE"`
}
