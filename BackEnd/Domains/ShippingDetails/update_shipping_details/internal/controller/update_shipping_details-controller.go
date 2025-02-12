package controller

import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	requests "update_shipping_details/internal/data/requests"
	responses "update_shipping_details/internal/data/responses"
	services "update_shipping_details/internal/service"
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