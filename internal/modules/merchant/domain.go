package merchant

import "github.com/xcurvnubaim/pbkk-golang/internal/modules/common"

// "gorm.io/gorm"

type MerchantModel struct {
	common.BaseModels
	Nama string `gorm:"name"`
    Email string `gorm:"email"`
    Alamat string `gorm:"alamat"`
}
func (MerchantModel) TableName() string {
	return "merchants"
}