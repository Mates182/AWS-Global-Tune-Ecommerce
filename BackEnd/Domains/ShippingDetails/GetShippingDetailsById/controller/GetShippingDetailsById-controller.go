package controller

import (
	requests "get-shipping-details-by-id/data/requests"
	responses "get-shipping-details-by-id/data/responses"
	"get-shipping-details-by-id/models"
	services "get-shipping-details-by-id/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetShippingDetailsByIdController struct {
	GetShippingDetailsByIdService services.GetShippingDetailsByIdService
}

func NewGetShippingDetailsByIdController(service services.GetShippingDetailsByIdService) *GetShippingDetailsByIdController {
	return &GetShippingDetailsByIdController{
		GetShippingDetailsByIdService: service,
	}
}

func (ctrl *GetShippingDetailsByIdController) GetShippingDetailsById(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.IndentedJSON(http.StatusBadRequest, responses.GetShippingDetailsByIdResponse{Message: "ID is required"})
		return
	}
	request := requests.GetShippingDetailsByIdRequest{ShippingDetails: models.ShippingDetails{ID: ID}}

	status, res := ctrl.GetShippingDetailsByIdService.GetShippingDetailsByIdHandler(request)

	c.IndentedJSON(status, res)
}
