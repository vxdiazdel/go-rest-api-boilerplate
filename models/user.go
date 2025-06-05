package models

type User struct {
	BaseModel
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"-" binding:"required,min=8"`
}
