// auto-generated with ginshot
package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	requests "create-tracking-details-service/data/requests"
	responses "create-tracking-details-service/data/responses"
	services "create-tracking-details-service/service"
)

type CreateTrackingDetailsController struct {
	CreateTrackingDetailsService services.CreateTrackingDetailsService
}

func NewCreateTrackingDetailsController(service services.CreateTrackingDetailsService) *CreateTrackingDetailsController {
	return &CreateTrackingDetailsController{
		CreateTrackingDetailsService: service,
	}
}

func (ctrl *CreateTrackingDetailsController) CreateTrackingDetails(c *gin.Context) {
	var request requests.CreateTrackingDetailsRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.CreateTrackingDetailsResponse{Message: "Invalid request body"})
		return
	}
	status, res := ctrl.CreateTrackingDetailsService.CreateTrackingDetailsHandler(request)

	c.IndentedJSON(status, res)
}
