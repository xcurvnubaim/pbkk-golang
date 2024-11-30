package interfaces

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/domain"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
)

type AuthRepository interface {
	RegisterUser(data *domain.RegisterUser) e.ApiError
	GetUserByEmail(email string) (*domain.User, e.ApiError)
}
