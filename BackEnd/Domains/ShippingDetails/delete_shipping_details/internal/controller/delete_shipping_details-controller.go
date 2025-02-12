package controller

import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	requests "delete_shipping_details/internal/data/requests"
	responses "delete_shipping_details/internal/data/responses"
	services "delete_shipping_details/internal/service"
)

type DeleteShippingDetailsController struct {
	DeleteShippingDetailsService services.DeleteShippingDetailsService
}

func NewDeleteShippingDetailsController(service services.DeleteShippingDetailsService) *DeleteShippingDetailsController {
	return &DeleteShippingDetailsController{
		DeleteShippingDetailsService: service,
	}
}

func (ctrl *DeleteShippingDetailsController) DeleteShippingDetails(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.IndentedJSON(http.StatusBadRequest, responses.DeleteShippingDetailsResponse{Message: "ID is required"})
		return
	}
	request := requests.DeleteShippingDetailsRequest{ID: ID}

	status, res := ctrl.DeleteShippingDetailsService.DeleteShippingDetailsHandler(request)

	c.IndentedJSON(status, res)
}