package common

import (
	"time"
	// "gorm.io/gorm"
)

type BaseModels struct {
	ID        int32     `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP, onUpdate:CURRENT_TIMESTAMP"`
	// DeletedAt gorm.DeletedAt
}
