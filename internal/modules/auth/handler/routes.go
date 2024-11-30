package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/modules/auth/interfaces"
)

type AuthHandler struct {
	authUseCase interfaces.AuthUseCase
	app         *gin.Engine
}

func NewAuthHandler(app *gin.Engine, authUseCase interfaces.AuthUseCase) {
	authHandler := &AuthHandler{
		app:         app,
		authUseCase: authUseCase,
	}

	authHandler.Routes()
}

func (ah *AuthHandler) Routes() {
	authentication := ah.app.Group("/auth")
	{
		authentication.GET("/register", ah.Register)
	}
}
