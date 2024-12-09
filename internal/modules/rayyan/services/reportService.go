// services/reportService.go
package services

import (
	"fmt"

	"github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/db"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/models"
)

// GenerateSalesReport untuk menghasilkan laporan berdasarkan tanggal
func GenerateSalesReport(startDate, endDate string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	// Query transaksi berdasarkan rentang tanggal
	result := db.DB.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&transactions)
	if result.Error != nil {
		return nil, fmt.Errorf("Error fetching sales report: %v", result.Error)
	}
	return transactions, nil
}
