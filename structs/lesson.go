package payload_struct

import "github.com/google/uuid"

type SCreateLessonPayload struct {
	Name       string    `json:"name"`
	AuthorId   uuid.UUID `json:"author_id"`
	CategoryID uuid.UUID `json:"category_id"`
}

type SGetLessonPayload struct {
	BranchId uuid.UUID `json:"branch"`
}
