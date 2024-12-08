package merchant

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
)

type IUseCase interface {
	CreateMerchant(request *CreateMerchantRequestDTO) (*CreateMerchantResponseDTO, e.ApiError)
	DeleteMerchant(id string) e.ApiError
	GetMerchant(id string) (*CreateMerchantResponseDTO, e.ApiError)
	GetAllMerchants() ([]CreateMerchantResponseDTO, e.ApiError)
	UpdateMerchant(id string, request *CreateMerchantEditRequestDTO) (*CreateMerchantResponseDTO, e.ApiError)
}

type useCase struct {
	repository IRepository
}

func NewuseCase(repository IRepository) *useCase {
	return &useCase{repository}
}

// func (u *useCase
func (u *useCase) CreateMerchant(request *CreateMerchantRequestDTO) (*CreateMerchantResponseDTO, e.ApiError) {

	MerchantModel := &MerchantModel{
		Nama: request.Nama,
		Email: request.Email,
		Alamat: request.Alamat,
	}

	MerchantModel, err := u.repository.CreateMerchant(MerchantModel)
	if err != nil {
		return nil, err
	}

	response := &CreateMerchantResponseDTO{
		ID:   MerchantModel.ID,
		Nama: request.Nama,
		Email: request.Email,
		Alamat: request.Alamat,
	}

	return response, nil
}

func (u *useCase) DeleteMerchant(id string) e.ApiError {
	// Memanggil repository untuk menghapus merchant
	err := u.repository.DeleteMerchant(id)
	if err != nil {
		return err
	}
	
	return nil
}

func (u *useCase) GetMerchant(id string) (*CreateMerchantResponseDTO, e.ApiError) {
	// Memanggil repository untuk mendapatkan merchant berdasarkan id
	merchant, err := u.repository.GetMerchant(id)
	if err != nil {
		return nil, err
	}

	// Mengembalikan response dengan data merchant
	response := &CreateMerchantResponseDTO{
		ID:    merchant.ID,
		Nama:  merchant.Nama,
		Email: merchant.Email,
		Alamat: merchant.Alamat,

	}

	return response, nil
}

func (u *useCase) GetAllMerchants() ([]CreateMerchantResponseDTO, e.ApiError) {
	merchants, err := u.repository.GetAllMerchants()
	if err != nil {
		return nil, err
	}

	var response []CreateMerchantResponseDTO
	for _, merchant := range merchants {
		response = append(response, CreateMerchantResponseDTO{
			ID:    merchant.ID,
			Nama:  merchant.Nama,
			Email: merchant.Email,
			Alamat: merchant.Alamat,
		})
	}

	return response, nil
}

func (u *useCase) UpdateMerchant(id string, request *CreateMerchantEditRequestDTO) (*CreateMerchantResponseDTO, e.ApiError) {
	merchant, err := u.repository.UpdateMerchant(id, request)
	if err != nil {
		return nil, err
	}

	response := &CreateMerchantResponseDTO{
		ID:    merchant.ID,
		Nama:  merchant.Nama,
		Email:  merchant.Email,
		Alamat:  merchant.Alamat,
	}

	return response, nil
}
