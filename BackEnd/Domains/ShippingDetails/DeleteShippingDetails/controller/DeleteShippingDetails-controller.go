package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	requests "delete-shipping-details/data/requests"
	responses "delete-shipping-details/data/responses" 
	services "delete-shipping-details/service"
	"delete-shipping-details/models"
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
	request := requests.DeleteShippingDetailsRequest{ShippingDetails: models.ShippingDetails{ID: ID}}

	status, res := ctrl.DeleteShippingDetailsService.DeleteShippingDetailsHandler(request)

	c.IndentedJSON(status, res)
}