package model

import (
	"time"
)

type SAuthentModel struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(50);not null" json:"username" binding:"required"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type SUserModel struct {
	ID        string        `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FullName  string        `gorm:"size:255;not null" json:"fullname"`
	Email     string        `gorm:"size:255;not null;unique" json:"email"`
	JobTitle  string        `gorm:"size:50" json:"job_title"`
	Country   string        `gorm:"size:50" json:"country"`
	AuthID    string        `gorm:"type:uuid;not null" json:"auth_id"`             // Foreign key to Authentication
	Auth      SAuthentModel `gorm:"foreignKey:AuthID;constraint:OnDelete:CASCADE"` // Reference to SAuthentModel
	CreatedAt time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
