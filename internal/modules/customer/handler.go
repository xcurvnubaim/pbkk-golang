package customer

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xcurvnubaim/pbkk-golang/internal/middleware"
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
	customer := h.app.Group(prefix).Use(middleware.AuthenticateJWT())
	{
		customer.GET("/", h.GetCustomer)
		customer.POST("/", h.CreateCustomer)
		customer.PUT("/:id", h.UpdateCustomer)
		customer.DELETE("/:id", h.DeleteCustomer)
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

func (h *Handler) GetCustomer(c *gin.Context) {
	res, apiErr := h.useCase.FindAllCustomer()

	if apiErr != nil {
		errMsg := apiErr.Error()
		c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to get customer", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User logged in successfully", &res))
}

func (h *Handler) UpdateCustomer(c *gin.Context) {
	var customer UpdateCustomerRequestDTO
	id := c.Param("id")

	// Convert string to int
    Id, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(400, app.NewErrorResponse("Bad Request", nil))
        return
    } 

	customer = UpdateCustomerRequestDTO{
		ID: int32(Id),
	}

	if err := c.ShouldBindJSON(&customer); err != nil {
		var errMessages = CustomValidator.FormatValidationErrors(err)
		c.JSON(400, app.NewErrorResponse("Validation Error", &errMessages))
		return
	}

	res, apiErr := h.useCase.UpdateCustomerById(&customer)

	if apiErr != nil {
		errMsg := apiErr.Error()
		c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to update customer", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("User logged in successfully", &res))
}

func (h *Handler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	// Convert string to int
    Id, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(400, app.NewErrorResponse("Bad Request", nil))
        return
    } 

	req := DeleteCustomerRequestDTO{
		ID: int32(Id),
	}

	data, apiErr := h.useCase.DeleteCustomerById(&req)

	if apiErr != nil {
		errMsg := apiErr.Error()
		c.JSON(apiErr.Code(), app.NewErrorResponse("Failed to delete customer", &errMsg))
		return
	}

	c.JSON(200, app.NewSuccessResponse("Customer Deleted successfully", data))
}