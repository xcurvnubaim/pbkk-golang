package dto

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}