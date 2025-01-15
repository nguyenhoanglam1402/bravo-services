package model

import (
	"time"

	"gorm.io/gorm"
)

type SLessonModel struct {
	ID       string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name     string `gorm:"column:name;type:varchar;size:255"`
	Author   string `gorm:"column:author;type:varchar;size:255"`
	RawData  string `gorm:"column:raw_data;type:varchar;size:255"`
	CompData string `gorm:"column:comp_data;type:varchar"`

	CreatedAt time.Time      `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"type:timestamp(3);not null"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp(3)"`
}
