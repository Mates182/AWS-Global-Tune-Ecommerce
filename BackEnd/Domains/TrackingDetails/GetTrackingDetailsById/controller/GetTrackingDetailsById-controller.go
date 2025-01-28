// auto-generated with ginshot
package controller

import (
	requests "get-tracking-details-by-id-service/data/requests"
	services "get-tracking-details-by-id-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetTrackingDetailsByIdController struct {
	GetTrackingDetailsByIdService services.GetTrackingDetailsByIdService
}

func NewGetTrackingDetailsByIdController(service services.GetTrackingDetailsByIdService) *GetTrackingDetailsByIdController {
	return &GetTrackingDetailsByIdController{
		GetTrackingDetailsByIdService: service,
	}
}

func (ctrl *GetTrackingDetailsByIdController) GetTrackingDetailsById(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is required"})
		return
	}
	request := requests.GetTrackingDetailsByIdRequest{OrderID: ID}
	status, res := ctrl.GetTrackingDetailsByIdService.GetTrackingDetailsByIdHandler(request)

	c.IndentedJSON(status, res)
}
