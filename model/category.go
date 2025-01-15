package model

import (
	"time"

	"gorm.io/gorm"
)

type SCategoryModel struct {
	ID   string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name string `gorm:"column:name;type:varchar;size:255"`

	CreatedAt time.Time      `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp(3);not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp(3)"`
}
