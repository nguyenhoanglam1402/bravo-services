package model

type Authentication struct {
	Id         string `json:"id"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}

type User struct {
	Id         string `json:"id"`
	Fullname   string `json:"fullname" binding:"required"`
	Email      string `json:"email" binding:"required"`
	JobTitle   string `json:"job_title" binding:"required"`
	Country    string `json:"country" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
