package payload_struct

type SCreateClassPayload struct {
	ClassName     string `json:"class_name" binding:"required"`
	EduGroupID    string `json:"edu_group_id" binding:"required"`
	EnrollmentKey string `json:"enrollment_key" binding:"required"`
}
