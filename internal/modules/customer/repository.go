package customer

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateCustomer(user *CustomerModel) (*CustomerModel, e.ApiError)
	FindAllCustomer() ([]CustomerModel, e.ApiError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateCustomer(customer *CustomerModel) (*CustomerModel, e.ApiError) {
	result := r.db.Create(customer)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_CREATE_CUSTOMER_REPOSITORY_FAILED, result.Error.Error())
	}

	return customer, nil
}

func (r *repository) FindAllCustomer() ([]CustomerModel, e.ApiError) {
	var customers []CustomerModel
	result := r.db.Find(&customers)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_GET_ALL_CUSTOMER_REPOSITORY_FAILED, result.Error.Error())
	}

	return customers, nil
}