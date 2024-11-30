package usecase

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/interfaces"
)

type authUseCase struct {
	authRepository interfaces.AuthRepository
}

func NewAuthUseCase(authRepository interfaces.AuthRepository) *authUseCase {
	return &authUseCase{authRepository}
}
