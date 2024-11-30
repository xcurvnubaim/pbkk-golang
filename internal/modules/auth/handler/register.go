package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/dto"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/common"
)


func (ah *AuthHandler) Register(c *gin.Context) {
	var authentication dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&authentication); err != nil {
		errMsg := err.Error()
		c.JSON(400, &common.Response[dto.RegisterUserResponse]{
			Status:  false,
			Message: "Bad Request",
			Data:    nil,
			Error:   &errMsg,
		})
		return
	}

	// Register User
	if err := ah.authUseCase.RegisterUser(&authentication); err != nil {
		errMsg := err.Error()
		c.JSON(err.Code(), &common.Response[dto.RegisterUserResponse]{
			Status:  false,
			Message: "Failed to register user",
			Data:    nil,
			Error:   &errMsg,
		})
		return
	}

	c.JSON(200, &common.Response[dto.RegisterUserResponse]{
		Status:  true,
		Message: "Register Success",
		Data: &dto.RegisterUserResponse{
			Email: authentication.Email,
		},
		Error: nil,
	},
	)
	return
}