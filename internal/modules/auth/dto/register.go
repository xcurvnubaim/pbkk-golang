package dto

type RegisterUserRequest struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type RegisterUserResponse struct {
	Email string `json:"email"`
}