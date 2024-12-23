package customer

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/common"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
)

type IUseCase interface {
	CreateCustomer(request *CreateCustomerRequestDTO) (*CreateCustomerResponseDTO, e.ApiError)
	FindAllCustomer() (*FindAllCustomerResponseDTO, e.ApiError)
	UpdateCustomerById(request *UpdateCustomerRequestDTO) (*UpdateCustomerResponseDTO, e.ApiError)
	DeleteCustomerById(request *DeleteCustomerRequestDTO) (*DeleteCustomerResponseDTO, e.ApiError)
}

type useCase struct {
	repository IRepository
}

func NewuseCase(repository IRepository) *useCase {
	return &useCase{repository}
}

// func (u *useCase
func (u *useCase) CreateCustomer(request *CreateCustomerRequestDTO) (*CreateCustomerResponseDTO, e.ApiError) {

	customerModel := &CustomerModel{
		Nama: request.Nama,
		Umur: request.Umur,
		Asal: request.Asal,
		NoHp: request.NoHp,
	}

	customerModel, err := u.repository.CreateCustomer(customerModel)
	if err != nil {
		return nil, err
	}

	response := &CreateCustomerResponseDTO{
		ID:   customerModel.ID,
		Nama: customerModel.Nama,
		Umur: customerModel.Umur,
		Asal: customerModel.Asal,
		NoHp: customerModel.NoHp,
	}

	return response, nil
}

func (u *useCase) FindAllCustomer() (*FindAllCustomerResponseDTO, e.ApiError) {
	customers, err := u.repository.FindAllCustomer()
	if err != nil {
		return nil, err
	}

	var response FindAllCustomerResponseDTO
	for _, customer := range customers {
		response.Customers = append(response.Customers, CustomerDTO{
			ID:   customer.ID,
			Nama: customer.Nama,
			Umur: customer.Umur,
			Asal: customer.Asal,
			NoHp: customer.NoHp,
		})
	}


	return &response, nil
}

func (u *useCase) UpdateCustomerById(request *UpdateCustomerRequestDTO) (*UpdateCustomerResponseDTO, e.ApiError) {
	customerModel, err := u.repository.FindCustomerById(request.ID)

	if err != nil {
		return nil, err
	}
	
	customerModel = &CustomerModel{
		BaseModels: common.BaseModels{
			ID: request.ID,
			CreatedAt: customerModel.CreatedAt,
		},
		Nama: request.Nama,
		Umur: request.Umur,
		Asal: request.Asal,
		NoHp: request.NoHp,
	}

	customerModel, err = u.repository.UpdateCustomerById(customerModel)
	if err != nil {
		return nil, err
	}

	response := &UpdateCustomerResponseDTO{
		CustomerDTO: CustomerDTO{
			ID:   customerModel.ID,
			Nama: customerModel.Nama,
			Umur: customerModel.Umur,
			Asal: customerModel.Asal,
			NoHp: customerModel.NoHp,
		},
	}

	return response, nil
}

func (u *useCase) DeleteCustomerById(request *DeleteCustomerRequestDTO) (*DeleteCustomerResponseDTO, e.ApiError) {
	err := u.repository.DeleteCustomerById(request.ID)
	if err != nil {
		return nil, err
	}

	response := &DeleteCustomerResponseDTO{
		ID: request.ID,
	}

	return response, nil
}