package merchant

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/app"
	CustomValidator "github.com/xcurvnubaim/pbkk-golang/internal/pkg/validator"
)

type Handler struct {
	useCase IUseCase
	app *gin.Engine
}

func NewHandler(app *gin.Engine, useCase IUseCase, prefixApi string) {
	handler := &Handler{
		app:     app,
		useCase: useCase,
	}

	handler.Routes(prefixApi)
}

func (h *Handler) Routes(prefix string) {
	merchant := h.app.Group(prefix)
	{
		merchant.POST("/", h.CreateMerchant)
		merchant.DELETE("/:id", h.DeleteMerchant)
		merchant.GET("/:id", h.GetMerchant)
		merchant.GET("/", h.GetMerchant)
		merchant.PUT("/:id", h.UpdateMerchant)
		// authentication.POST("/register", h.Register)
		// authentication.POST("/login", h.Login)

		// authentication.Use(middleware.AuthenticateJWT())
		// {
		// 	authentication.GET("/me", h.GetMe)
		// }
	}
}

func (h *Handler) CreateMerchant(c *gin.Context) {
	var merchant CreateMerchantRequestDTO

	if err := c.ShouldBindJSON(&merchant); err != nil {
		var errMessages = CustomValidator.FormatValidationErrors(err)
		c.JSON(400, app.NewErrorResponse("Validation Error", &errMessages))
		return
	}

	res, apiErr := h.useCase.CreateMerchant(&merchant)

	if apiErr != nil {
		errMsg := apiErr.Error()
		c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to create merchant", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("Merchant Created successfully", &res))
}

func (h *Handler) DeleteMerchant(c *gin.Context) {
    id := c.Param("id") // Ambil ID dari parameter URL

    apiErr := h.useCase.DeleteMerchant(id)

    if apiErr != nil {
        errMsg := apiErr.Error()
        c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to delete merchant", &errMsg))
        return
    }

	response := &DeleteMerchantResponseDTO{
		Message: fmt.Sprintf("Merchant with ID %s deleted successfully", id),
	}
    c.JSON(200, app.NewSuccessResponse("Merchant deleted successfully", response))
}

func (h *Handler) GetMerchant(c *gin.Context) {
	id := c.Param("id") // Ambil ID dari parameter URL

	if id == "" {
		// Jika ID tidak diberikan, ambil semua merchants
		merchants, apiErr := h.useCase.GetAllMerchants()
		if apiErr != nil {
			errMsg := apiErr.Error()
			c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to fetch merchants", &errMsg))
			return
		}

		c.JSON(200, app.NewSuccessResponse("Merchants fetched successfully", &merchants))
	} else {
		// Jika ID diberikan, ambil merchant berdasarkan ID
		merchant, apiErr := h.useCase.GetMerchant(id)
		if apiErr != nil {
			errMsg := apiErr.Error()
			c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to fetch merchant", &errMsg))
			return
		}

		c.JSON(200, app.NewSuccessResponse("Merchant fetched successfully", &merchant))
	}
}

func (h *Handler) UpdateMerchant(c *gin.Context) {
	id := c.Param("id")
	var request CreateMerchantEditRequestDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		var errMessages = CustomValidator.FormatValidationErrors(err)
		c.JSON(400, app.NewErrorResponse("Validation Error", &errMessages))
		return
	}

	merchant, apiErr := h.useCase.UpdateMerchant(id, &request)
	if apiErr != nil {
		errMsg := apiErr.Error()
		c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to update merchant", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("Merchant updated successfully", merchant))
}
