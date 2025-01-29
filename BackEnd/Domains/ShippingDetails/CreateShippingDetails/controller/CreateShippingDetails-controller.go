package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	requests "create-shipping-details/data/requests"
	responses "create-shipping-details/data/responses" 
	services "create-shipping-details/service"
	
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