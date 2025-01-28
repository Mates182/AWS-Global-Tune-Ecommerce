package controller

import (
	requests "list-tracking-details/data/requests"
	responses "list-tracking-details/data/responses"
	services "list-tracking-details/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListTrackingDetailsController struct {
	ListTrackingDetailsService services.ListTrackingDetailsService
}

func NewListTrackingDetailsController(service services.ListTrackingDetailsService) *ListTrackingDetailsController {
	return &ListTrackingDetailsController{
		ListTrackingDetailsService: service,
	}
}

func (ctrl *ListTrackingDetailsController) ListTrackingDetails(c *gin.Context) {
	var request requests.ListTrackingDetailsRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ListTrackingDetailsResponse{Message: "Invalid request body"})
		return
	}

	status, res := ctrl.ListTrackingDetailsService.ListTrackingDetailsHandler(request)

	c.IndentedJSON(status, res)
}
