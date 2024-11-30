package domain

import "github.com/xcurvnubaim/pbkk-golang/internal/modules/common"

type User struct {
	common.BaseModels
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"default:'user'"`
}

func (User) TableName() string {
	return "users"
}