package controller

import (
	"net/http"
	"update_product/internal/auth"
	requests "update_product/internal/data/requests"
	responses "update_product/internal/data/responses"
	services "update_product/internal/service"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	UpdateProductService services.UpdateProductService
}

func NewUpdateProductController(service services.UpdateProductService) *UpdateProductController {
	return &UpdateProductController{
		UpdateProductService: service,
	}
}

func (ctrl *UpdateProductController) UpdateProduct(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid token"})
		return
	}
	if !auth.ValidateTokenMiddleware("admin", token) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid role or token"})
		return

	}
	var request requests.UpdateProductRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.UpdateProductResponse{Message: "Invalid request body"})
		return
	}
	status, res := ctrl.UpdateProductService.UpdateProductHandler(request)

	c.IndentedJSON(status, res)
}
