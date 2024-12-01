package auth

import (
	"github.com/google/uuid"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/common"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	RegisterUser(*RegisterUserDomain) e.ApiError
	GetUserByUsername(string) (*UserModel, e.ApiError)
	GetUserByID(uuid.UUID) (*UserModel, e.ApiError)
	GetAllUser() ([]UserModel, e.ApiError)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) RegisterUser(data *RegisterUserDomain) e.ApiError {
	user := &UserModel{
		BaseModels: common.BaseModels{
			ID: data.Id, // or however you assign the ID
		},
		Username: data.Username,
		Password: data.Password,
	}

	result := r.db.Create(user)
	if result.Error != nil {
		return e.NewApiError(e.ERROR_REGISTER_REPOSITORY_FAILED, result.Error.Error())
	}

	return nil
}

func (r *authRepository) GetUserByUsername(username string) (*UserModel, e.ApiError) {
	user := &UserModel{}
	result := r.db.Where("username = ?", username).First(user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, e.NewApiError(e.ERROR_USER_NOT_FOUND, "User not found")
		}
		return nil, e.NewApiError(e.ERROR_GET_USER_BY_EMAIL_REPOSITORY_FAILED, result.Error.Error())
	}

	return user, nil
}

func (r *authRepository) GetUserByID(id uuid.UUID) (*UserModel, e.ApiError) {
	user := &UserModel{}
	result := r.db.Where("id = ?", id).First(user)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_GET_USER_BY_ID_REPOSITORY_FAILED, result.Error.Error())
	}

	return user, nil
}

func (r *authRepository) GetAllUser() ([]UserModel, e.ApiError) {
	var users []UserModel
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_GET_ALL_USER_REPOSITORY_FAILED, result.Error.Error())
	}

	return users, nil
}
