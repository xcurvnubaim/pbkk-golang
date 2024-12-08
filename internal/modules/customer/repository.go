package customer

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateCustomer(user *CustomerModel) (*CustomerModel, e.ApiError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCustomer(user *CustomerModel) (*CustomerModel, e.ApiError) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_CREATE_CUSTOMER_REPOSITORY_FAILED, result.Error.Error())
	}

	return user, nil
}
