package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	requests "update-shipping-details/data/requests"
	responses "update-shipping-details/data/responses" 
	services "update-shipping-details/service"
	
)

type UpdateShippingDetailsController struct {
	UpdateShippingDetailsService services.UpdateShippingDetailsService
}

func NewUpdateShippingDetailsController(service services.UpdateShippingDetailsService) *UpdateShippingDetailsController {
	return &UpdateShippingDetailsController{
		UpdateShippingDetailsService: service,
	}
}

func (ctrl *UpdateShippingDetailsController) UpdateShippingDetails(c *gin.Context) {
	var request requests.UpdateShippingDetailsRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.UpdateShippingDetailsResponse{Message: "Invalid request body"})
		return
		}

	status, res := ctrl.UpdateShippingDetailsService.UpdateShippingDetailsHandler(request)

	c.IndentedJSON(status, res)
}