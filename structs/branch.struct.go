package payload_struct

import (
	"bravo-service/api/model"

	"github.com/google/uuid"
)

type SBranchPayload struct {
	VesionID   uuid.UUID `json:"current_version_id" binding:"required"`
	BranchName string    `json:"branch_name" binding:"required"`
	UserID     uuid.UUID `json:"user_id" binding:"required"`
}

type SCommitPayload struct {
	BranchID uuid.UUID `json:"branch_id" binding:"required"`
	AuthorID uuid.UUID `json:"author_id" binding:"required"`
	RawData  string    `json:"raw_data" binding:"required"`
	Message  string    `json:"message"`
	CompData string    `json:"comp_data"`
}

type SMergeBranchPayload struct {
	BranchID uuid.UUID `json:"branch_id" binding:"required"`
}

type SCommitBranchPayload struct {
	BranchID        uuid.UUID  `json:"branch_id" binding:"required"`
	MessageContent  string     `json:"message_content" binding:"required"`
	RawData         string     `json:"raw_data" binding:"required"`
	CompData        string     `json:"comp_data" binding:"required"`
	ParentVersionID *uuid.UUID `json:"parent_version_id"`
}

type SVersionBranchTreePayload struct {
	BranchID   uuid.UUID         `json:"branch_id" binding:"required"`
	BranchName string            `json:"branch_name" binding:"required"`
	Versions   *[]model.SVersion `json:"versions"`
}
