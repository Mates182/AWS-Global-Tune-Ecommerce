package controller

import (
	"delete_product/internal/auth"
	requests "delete_product/internal/data/requests"
	responses "delete_product/internal/data/responses"
	services "delete_product/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductController struct {
	DeleteProductService services.DeleteProductService
}

func NewDeleteProductController(service services.DeleteProductService) *DeleteProductController {
	return &DeleteProductController{
		DeleteProductService: service,
	}
}

func (ctrl *DeleteProductController) DeleteProduct(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Missing or invalid token"})
		return
	}
	if !auth.ValidateTokenMiddleware("admin", token) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid role or token"})
		return

	}
	ID := c.Param("id")
	if ID == "" {
		c.IndentedJSON(http.StatusBadRequest, responses.DeleteProductResponse{Message: "ID is required"})
		return
	}
	request := requests.DeleteProductRequest{ID: ID}

	status, res := ctrl.DeleteProductService.DeleteProductHandler(request)

	c.IndentedJSON(status, res)
}
