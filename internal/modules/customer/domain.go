package customer

import "github.com/xcurvnubaim/pbkk-golang/internal/modules/common"

// "gorm.io/gorm"

type CustomerModel struct {
	common.BaseModels
	Nama string `gorm:"column:nama"`
	Umur int32 `gorm:"column:umur"`
	Asal string `gorm:"column:asal"`
	NoHp string `gorm:"column:no_hp"`
}

func (CustomerModel) TableName() string {
	return "customers"
}
