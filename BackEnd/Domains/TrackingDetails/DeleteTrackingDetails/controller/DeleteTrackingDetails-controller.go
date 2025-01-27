// auto-generated with ginshot
package controller

import (
	requests "delete-tracking-details-service/data/requests"
	"delete-tracking-details-service/models"
	services "delete-tracking-details-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteTrackingDetailsController struct {
	DeleteTrackingDetailsService services.DeleteTrackingDetailsService
}

func NewDeleteTrackingDetailsController(service services.DeleteTrackingDetailsService) *DeleteTrackingDetailsController {
	return &DeleteTrackingDetailsController{
		DeleteTrackingDetailsService: service,
	}
}

func (ctrl *DeleteTrackingDetailsController) DeleteTrackingDetails(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is required"})
		return
	}
	request := requests.DeleteTrackingDetailsRequest{TrackingDetails: models.TrackingDetails{OrderID: ID}}
	status, res := ctrl.DeleteTrackingDetailsService.DeleteTrackingDetailsHandler(request)

	c.IndentedJSON(status, res)
}
