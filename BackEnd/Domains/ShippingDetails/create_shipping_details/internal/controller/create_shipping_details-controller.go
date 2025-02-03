package controller

import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	requests "create_shipping_details/internal/data/requests"
	responses "create_shipping_details/internal/data/responses"
	services "create_shipping_details/internal/service"
)

type CreateShippingDetailsController struct {
	CreateShippingDetailsService services.CreateShippingDetailsService
}

func NewCreateShippingDetailsController(service services.CreateShippingDetailsService) *CreateShippingDetailsController {
	return &CreateShippingDetailsController{
		CreateShippingDetailsService: service,
	}
}

func (ctrl *CreateShippingDetailsController) CreateShippingDetails(c *gin.Context) {
	var request requests.CreateShippingDetailsRequest
		if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.CreateShippingDetailsResponse{Message: "Invalid request body"})
		return
		}
	status, res := ctrl.CreateShippingDetailsService.CreateShippingDetailsHandler(request)

	c.IndentedJSON(status, res)
}