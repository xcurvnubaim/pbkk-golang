package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/file-uploads/interfaces"
)

type FileHandler struct {
	fileUseCase interfaces.FileUseCase
	app         *gin.Engine
}

func NewFileHandler(app *gin.Engine, fileUseCase interfaces.FileUseCase) {
	authHandler := &FileHandler{
		app:         app,
		fileUseCase: fileUseCase,
	}

	authHandler.Routes()
}

func (h *FileHandler) Routes() {
	file := h.app.Group("/file")
	{
		// authentication.GET("/register", ah.Register)
		file.POST("/upload", h.Upload)
	}
}
