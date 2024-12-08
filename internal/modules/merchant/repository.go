package merchant

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/e"
	"gorm.io/gorm"
)

type IRepository interface {
	CreateMerchant(user *MerchantModel) (*MerchantModel, e.ApiError)
	DeleteMerchant(id string) e.ApiError
	GetMerchant(id string) (*MerchantModel, e.ApiError)
	GetAllMerchants() ([]MerchantModel, e.ApiError)
	UpdateMerchant(id string, merchant *CreateMerchantEditRequestDTO) (*MerchantModel, e.ApiError)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateMerchant(user *MerchantModel) (*MerchantModel, e.ApiError) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_CREATE_Merchant_REPOSITORY_FAILED, result.Error.Error())
	}

	return user, nil
}

func (r *repository) DeleteMerchant(id string) e.ApiError {
    var merchant MerchantModel
    // Cari merchant berdasarkan ID
    result := r.db.Where("id = ?", id).First(&merchant)
    if result.Error != nil {
        return e.NewApiError(e.ERROR_DELETE_Merchant_REPOSITORY_FAILED, "Merchant not found")
    }

    // Hapus merchant
    deleteResult := r.db.Delete(&merchant)
    if deleteResult.Error != nil {
        return e.NewApiError(e.ERROR_DELETE_Merchant_REPOSITORY_FAILED, deleteResult.Error.Error())
    }

    return nil
}

func (r *repository) GetMerchant(id string) (*MerchantModel, e.ApiError) {
    var merchant MerchantModel
    // Cari merchant berdasarkan ID
    result := r.db.Where("id = ?", id).First(&merchant)
    if result.Error != nil {
        return nil, e.NewApiError(e.ERROR_GET_MERCHANT_REPOSITORY_FAILED, "Merchant not found")
    }

    return &merchant, nil
}

func (r *repository) GetAllMerchants() ([]MerchantModel, e.ApiError) {
	var merchants []MerchantModel
	result := r.db.Find(&merchants)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_GET_MERCHANT_REPOSITORY_FAILED, result.Error.Error())
	}

	return merchants, nil
}

func (r *repository) UpdateMerchant(id string, merchant *CreateMerchantEditRequestDTO) (*MerchantModel, e.ApiError) {
	var existingMerchant MerchantModel
	result := r.db.Where("id = ?", id).First(&existingMerchant)
	if result.Error != nil {
		return nil, e.NewApiError(e.ERROR_MERCHANT_NOT_FOUND, "Merchant not found")
	}

	// Update fields dynamically based on the request
	if merchant.Nama != "" {
		existingMerchant.Nama = merchant.Nama
	}
	if merchant.Email != "" {
		existingMerchant.Email = merchant.Email
	}
	if merchant.Alamat != "" {
		existingMerchant.Alamat = merchant.Alamat
	}


	// Save the updated merchant data
	r.db.Save(&existingMerchant)

	return &existingMerchant, nil
}
