package models

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,max=100"`
	Password string `json:"password" binding:"required,min=8"`
}
