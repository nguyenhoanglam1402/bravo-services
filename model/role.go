package model

import (
	"time"

	"gorm.io/gorm"
)

type SRoleModel struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time      `gorm:"type:timestamp(3);not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"type:timestamp(3);not null" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp(3)" json:"-"`
}
