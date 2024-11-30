package repository

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/domain"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) RegisterUser(data *domain.RegisterUser) e.ApiError {
	user := &domain.User{
		Email:    data.Email,
		Password: data.Password,
	}

	result := r.db.Create(user)
	if result.Error != nil {
		return e.NewApiError(10011, result.Error.Error())
	}

	return nil
}

func (r *authRepository) GetUserByEmail(email string) (*domain.User, e.ApiError) {
	user := &domain.User{}
	result := r.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, e.NewApiError(10012, result.Error.Error())
	}

	return user, nil
}
