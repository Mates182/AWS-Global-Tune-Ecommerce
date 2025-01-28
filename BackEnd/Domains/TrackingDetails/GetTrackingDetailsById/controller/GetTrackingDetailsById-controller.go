// auto-generated with ginshot
package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	requests "get-tracking-details-by-id-service/data/requests"
	responses "get-tracking-details-by-id-service/data/responses"
	services "get-tracking-details-by-id-service/service"
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
	var request requests.GetTrackingDetailsByIdRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.GetTrackingDetailsByIdResponse{Message: "Invalid request body"})
		return
	}
	status, res := ctrl.GetTrackingDetailsByIdService.GetTrackingDetailsByIdHandler(request)

	c.IndentedJSON(status, res)
}
