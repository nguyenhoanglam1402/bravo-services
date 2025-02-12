package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SVersion struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	BranchID        uuid.UUID  `gorm:"type:uuid;not null"`
	AuthorID        uuid.UUID  `gorm:"type:uuid;not null"`
	ParentVersionID *uuid.UUID `gorm:"type:uuid"` // Nullable for initial versions
	RawData         string     `gorm:"type:text;not null"`
	CompData        string     `gorm:"type:text"`
	CreatedAt       time.Time  `gorm:"type:timestamp(3);default:CURRENT_TIMESTAMP"`

	// Relationships
	Branch        *SBranch    `gorm:"foreignKey:BranchID;references:ID;constraint:OnDelete:CASCADE"`
	Author        *SUserModel `gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:CASCADE"`
	ParentVersion *SVersion   `gorm:"foreignKey:ParentVersionID;references:ID;constraint:OnDelete:SET NULL"`
}

// BeforeCreate sets a new UUID before inserting a record
func (v *SVersion) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}
