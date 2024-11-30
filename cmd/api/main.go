package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/configs"
	"github.com/xcurvnubaim/pbkk-golang/internal/database"

	AuthHandler "github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/handler"
	AuthInterfaces "github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/interfaces"
	AuthRepository "github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/repository"
	AuthUseCase "github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/usecase"

	FileHandler "github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/handler"
	FileInterfaces "github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/interfaces"
	FileRepository "github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/repository"
	FileUseCase "github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/usecase"
)

func main() {
	// Setup configuration
	if err := configs.Setup(); err != nil {
		panic(err)
	}

	// Setup for production
	if configs.Config.ENV_MODE == "production" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("Production mode")
	}

	// Start the server
	r := gin.Default()

	// Setup Database
	db := database.New()

	// Module Authentication
	var authRepository AuthInterfaces.AuthRepository = AuthRepository.NewAuthRepository(db)
	var authUseCase AuthInterfaces.AuthUseCase = AuthUseCase.NewAuthUseCase(authRepository)
	AuthHandler.NewAuthHandler(r, authUseCase)

	// Module File Uploads
	var fileRepository FileInterfaces.FileRepository = FileRepository.NewFileRepository(db)
	var fileUseCase FileInterfaces.FileUseCase = FileUseCase.NewFileUseCase(fileRepository)
	FileHandler.NewFileHandler(r, fileUseCase)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":" + configs.Config.APP_PORT); err != nil {
		panic(err)
	}
}
