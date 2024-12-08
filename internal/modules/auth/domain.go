package auth

import (
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/common"
)

type (
	RegisterUserDomain struct {
		Id       int32
		Username string
		Password string
	}

	UserModel struct {
		common.BaseModels
		Username string `gorm:"unique;not null"`
		Password string `gorm:"not null"`
		Role     string `gorm:"default:'user'"`
	}

	PayloadToken struct {
		ID   int32
		Role string
	}
)

func (UserModel) TableName() string {
	return "users"
}
