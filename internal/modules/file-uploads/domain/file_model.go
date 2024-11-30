package domain

import "github.com/xcurvnubaim/pbkk-golang/internal/modules/common"

type File struct {
	common.BaseModels
	FileName string `gorm:"not null"`
	FilePath string `gorm:"not null"`
}

func (File) TableName() string {
	return "file_uploads"
}
