package payload_struct

type SLoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SSignUpPayload struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	JobTitle string `json:"job_title" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Country  string `json:"country" binding:"required"`
	Password string `json:"password" binding:"required"`
}
