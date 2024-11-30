package interfaces

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/dto"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
)

type AuthUseCase interface {
	RegisterUser(data *dto.RegisterUserRequest) e.ApiError
}
