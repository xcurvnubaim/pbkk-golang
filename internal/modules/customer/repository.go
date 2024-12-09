package customer

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateCustomer(user *CustomerModel) (*CustomerModel, e.ApiError)
	FindAllCustomer() ([]CustomerModel, e.ApiError)
	UpdateCustomerById(user *CustomerModel) (*CustomerModel, e.ApiError)
	DeleteCustomerById(id int32) e.ApiError
	FindCustomerById(id int32) (*CustomerModel, e.ApiError)
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

func (r *repository) FindCustomerById(id int32) (*CustomerModel, e.ApiError) {
	var customer CustomerModel
	result := r.db.First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, e.NewApiError(e.ERROR_CUSTOMER_NOT_FOUND, "customer not found")
		}
		return nil, e.NewApiError(e.ERROR_GET_CUSTOMER_BY_ID_REPOSITORY_FAILED, result.Error.Error())
	}

	return &customer, nil
}

func (r *repository) UpdateCustomerById(customer *CustomerModel) (*CustomerModel, e.ApiError) {
	result := r.db.Save(customer)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_UPDATE_CUSTOMER_REPOSITORY_FAILED, result.Error.Error())
	}

	return customer, nil
}

func (r *repository) DeleteCustomerById(id int32) e.ApiError {
	result := r.db.Delete(&CustomerModel{}, id)
	if result.Error != nil {
		return e.NewApiError(e.ERROR_DELETE_CUSTOMER_REPOSITORY_FAILED, result.Error.Error())
	}

	return nil
}