package controller

import (
	"create_product/internal/auth"
	requests "create_product/internal/data/requests"
	responses "create_product/internal/data/responses"
	services "create_product/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	CreateProductService services.CreateProductService
}

func NewCreateProductController(service services.CreateProductService) *CreateProductController {
	return &CreateProductController{
		CreateProductService: service,
	}
}

func (ctrl *CreateProductController) CreateProduct(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid token"})
		return
	}
	if !auth.ValidateTokenMiddleware("admin", token) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid role or token"})
		return

	}

	var request requests.CreateProductRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.CreateProductResponse{Message: "Invalid request body"})
		return
	}
	status, res := ctrl.CreateProductService.CreateProductHandler(request)

	c.IndentedJSON(status, res)
}
