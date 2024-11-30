package usecase

import (
	"fmt"
	"log"

	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/domain"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/dto"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
)

func (uc *authUseCase) RegisterUser(data *dto.RegisterUserRequest) e.ApiError {
	userCheck, _ := uc.authRepository.GetUserByEmail(data.Email)

	if userCheck != nil {
		return e.NewApiError(400, "Email already registered")
	}

	user := &domain.RegisterUser{
		Email:    data.Email,
		Password: data.Password,
	}

	if err := uc.authRepository.RegisterUser(user); err != nil {
		log.Println(err.Error())
		return e.NewApiError(500, fmt.Sprintf("Internal Server Error (%d)", err.Code()))
	}
	return nil
}
