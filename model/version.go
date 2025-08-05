package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SVersion struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	BranchID        uuid.UUID  `gorm:"type:uuid;not null" json:"branch_id"`
	AuthorID        uuid.UUID  `gorm:"type:uuid;not null" json:"author_id"`
	ParentVersionID *uuid.UUID `gorm:"type:uuid" json:"parent_version_id"` // Nullable for initial versions
	RawData         string     `gorm:"type:text;not null" json:"raw_data"`
	CompData        string     `gorm:"type:text" json:"comp_data"`
	Message         string     `gorm:"type:varchar" json:"message"`
	CreatedAt       time.Time  `gorm:"type:timestamp(3);default:CURRENT_TIMESTAMP" json:"created_at"`

	// Relationships
	Branch        *SBranch    `gorm:"foreignKey:BranchID;references:ID;constraint:OnDelete:CASCADE" json:"branch"`
	Author        *SUserModel `gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:CASCADE" json:"author"`
	ParentVersion *SVersion   `gorm:"foreignKey:ParentVersionID;references:ID;constraint:OnDelete:SET NULL" json:"parent_version"`
}

// BeforeCreate sets a new UUID before inserting a record
func (v *SVersion) BeforeCreate(tx *gorm.DB) (err error) {
	if v.ID == uuid.Nil {
		v.ID = uuid.New()
	}
	return
}
