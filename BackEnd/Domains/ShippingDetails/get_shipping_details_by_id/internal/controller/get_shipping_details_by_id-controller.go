package controller

import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	requests "get_shipping_details_by_id/internal/data/requests"
	responses "get_shipping_details_by_id/internal/data/responses"
	services "get_shipping_details_by_id/internal/service"
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
	request := requests.GetShippingDetailsByIdRequest{ID: ID}

	status, res := ctrl.GetShippingDetailsByIdService.GetShippingDetailsByIdHandler(request)

	c.IndentedJSON(status, res)
}