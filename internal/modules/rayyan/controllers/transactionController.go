package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/db"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/models"
	"gorm.io/gorm"
)

// GetTransactions handles fetching all transactions
func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	if err := db.DB.
		Order("created_at desc").
		Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransactionByID handles fetching a single transaction
func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	if err := db.DB.
		First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// CreateTransaction handles creating a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Decrement product stock
	for _, item := range transaction.Items {
		if err := db.DB.Model(&models.Product{}).
			Where("id = ? AND stock_quantity >= ?", item.ProductID, item.Quantity).
			UpdateColumn("stock_quantity", gorm.Expr("stock_quantity - ?", item.Quantity)).Error; err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
			return
		}
	}
		

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction created successfully",
		"transaction": transaction,
	})
}
