package repository

import (
	"gorm.io/gorm"
)

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *fileRepository {
	return &fileRepository{db}
}
