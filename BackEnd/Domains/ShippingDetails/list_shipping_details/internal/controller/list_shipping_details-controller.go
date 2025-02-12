package controller

import (
	
	"github.com/gin-gonic/gin"
	
	services "list_shipping_details/internal/service"
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
	
	status, res := ctrl.ListShippingDetailsService.ListShippingDetailsHandler()

	c.IndentedJSON(status, res)
}