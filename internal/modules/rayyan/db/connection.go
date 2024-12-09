// db/connection.go
package db

import (
	"fmt"
	"log"

	"github.com/xcurvnubaim/pbkk-golang/internal/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection
func Init() {
	dsn := configs.Config.DatabaseURL + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to the database.")
}
