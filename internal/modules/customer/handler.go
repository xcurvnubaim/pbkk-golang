package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/pkg/app"
	CustomValidator "github.com/xcurvnubaim/pbkk-golang/internal/pkg/validator"
)

type Handler struct {
	useCase IUseCase
	app     *gin.Engine
}

func NewHandler(app *gin.Engine, useCase IUseCase, prefixApi string) {
	handler := &Handler{
		app:     app,
		useCase: useCase,
	}

	handler.Routes(prefixApi)
}

func (h *Handler) Routes(prefix string) {
	customer := h.app.Group(prefix)
	{
		customer.POST("/", h.CreateCustomer)
		// authentication.POST("/register", h.Register)
		// authentication.POST("/login", h.Login)

		// authentication.Use(middleware.AuthenticateJWT())
		// {
		// 	authentication.GET("/me", h.GetMe)
		// }
	}
}

func (h *Handler) CreateCustomer(c *gin.Context) {
	var customer CreateCustomerRequestDTO

	if err := c.ShouldBindJSON(&customer); err != nil {
		var errMessages = CustomValidator.FormatValidationErrors(err)
		c.JSON(400, app.NewErrorResponse("Validation Error", &errMessages))
		return
	}

	res, apiErr := h.useCase.CreateCustomer(&customer)

	if apiErr != nil {
		errMsg := apiErr.Error()
		c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to create customer", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User logged in successfully", &res))
}
