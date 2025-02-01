package controller

import (
	requests "list-shipping-details/data/requests"
	services "list-shipping-details/service"

	"github.com/gin-gonic/gin"
)

type ListShippingDetailsController struct {
	ListShippingDetailsService services.ListShippingDetailsService
}

func NewListShippingDetailsController(service services.ListShippingDetailsService) *ListShippingDetailsController {
	return &ListShippingDetailsController{
		ListShippingDetailsService: service,
	}
}

func (ctrl *ListShippingDetailsController) ListShippingDetails(c *gin.Context) {
	var request requests.ListShippingDetailsRequest

	status, res := ctrl.ListShippingDetailsService.ListShippingDetailsHandler(request)

	c.IndentedJSON(status, res)
}
