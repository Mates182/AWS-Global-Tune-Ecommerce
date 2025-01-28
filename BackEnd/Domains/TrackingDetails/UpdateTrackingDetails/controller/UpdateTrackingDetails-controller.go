// auto-generated with ginshot
package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	requests "update-tracking-details-service/data/requests"
	responses "update-tracking-details-service/data/responses"
	services "update-tracking-details-service/service"
)

type UpdateTrackingDetailsController struct {
	UpdateTrackingDetailsService services.UpdateTrackingDetailsService
}

func NewUpdateTrackingDetailsController(service services.UpdateTrackingDetailsService) *UpdateTrackingDetailsController {
	return &UpdateTrackingDetailsController{
		UpdateTrackingDetailsService: service,
	}
}

func (ctrl *UpdateTrackingDetailsController) UpdateTrackingDetails(c *gin.Context) {
	var request requests.UpdateTrackingDetailsRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.UpdateTrackingDetailsResponse{Message: "Invalid request body"})
		return
	}
	status, res := ctrl.UpdateTrackingDetailsService.UpdateTrackingDetailsHandler(request)

	c.IndentedJSON(status, res)
}
