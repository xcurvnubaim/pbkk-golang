package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/configs"
	"github.com/xcurvnubaim/pbkk-golang/internal/database"
	"github.com/xcurvnubaim/pbkk-golang/internal/middleware"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/customer"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/merchant"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/controllers"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/utils"
	db2 "github.com/xcurvnubaim/pbkk-golang/internal/modules/rayyan/db" 
)

func main() {
	// Setup configuration
	if err := configs.Setup(".env"); err != nil {
		panic(err)
	}

	// Setup for production
	if configs.Config.ENV_MODE == "production" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Production mode")
	}

	// Start the server
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	// Setup Database
	db := database.New()
	
	var authRepository auth.IAuthRepository = auth.NewAuthRepository(db)
	var authService auth.IAuthUseCase = auth.NewAuthUseCase(authRepository)
	auth.NewAuthHandler(r, authService, "/api/v1/auth")

	var customerRepository customer.IRepository = customer.NewRepository(db)
	var customerService customer.IUseCase = customer.NewuseCase(customerRepository)
	customer.NewHandler(r, customerService, "/api/v1/customer")

	var merchantRepository merchant.IRepository = merchant.NewRepository(db)
	var merchantService merchant.IUseCase = merchant.NewuseCase(merchantRepository)
	merchant.NewHandler(r, merchantService, "/api/v1/merchant")

	db2.DB = db

	auth := r.Group("/api/v2")
	auth.Use(utils.AuthMiddleware()) // Middleware for JWT token validation
	{
		// Sales Report Route (protected)
		auth.GET("/sales-report", controllers.GetSalesReport)

		// auth.POST("/users", controllers.Register)
		// auth.GET("/users", controllers.GetAllUsers)
		// auth.PUT("/users/:id", controllers.UpdateUser)
		// auth.DELETE("/users/:id", controllers.DeleteUser)
		// Product Routes	
		auth.POST("/products", controllers.CreateProduct) // Create product
		auth.GET("/products", controllers.GetProducts)    // Get all products
		// auth.GET("/products/:id", controllers.GetTransactionByID) // Get product by ID
		auth.PUT("/products/:id", controllers.UpdateProduct)    // Update product
		auth.DELETE("/products/:id", controllers.DeleteProduct) // Delete Product by ID

		// Stock Routes
		auth.POST("/stocks", controllers.CreateStock) // Create stock for product

		// Transaction Routes
		auth.GET("/transactions", controllers.GetTransactions)
		auth.GET("/transactions/:id", controllers.GetTransactionByID)
		auth.POST("/transactions", controllers.CreateTransaction)
		// auth.GET("/transactions", controllers.GetTransactions)    // Get All Transactions
	}

	// Serve static files
	// r.Static("/static", "./static") // Serve files with prefix `/static`

	// Dynamically map index.html files to routes
	staticDir := "./static"
	mapStaticRoutes(r, staticDir)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong test",
		})
	})

	if err := r.Run(":" + configs.Config.APP_PORT); err != nil {
		panic(err)
	}
}

// mapStaticRoutes dynamically maps routes for index.html files
func mapStaticRoutes(router *gin.Engine, baseDir string) {
	fmt.Printf("masuk mapstatic route")
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-index.html files
		if info.IsDir() || !strings.HasSuffix(info.Name(), "index.html") {
			return nil
		}

		// Generate route path from file path
		routePath := strings.TrimPrefix(path, baseDir)
		routePath = strings.TrimSuffix(routePath, "index.html")

		// Normalize path separators for cross-platform compatibility
		routePath = filepath.ToSlash(routePath)

		// Remove the "static" prefix if present
		routePath = strings.TrimPrefix(routePath, "static/")

		if routePath == "" {
			routePath = "/"
		}

		// Serve the index.html file at the generated route
		router.StaticFile(routePath, path)
		fmt.Printf("Mapped %s to %s\n", routePath, path)

		return nil
	})

	if err != nil {
		panic(err)
	}
}