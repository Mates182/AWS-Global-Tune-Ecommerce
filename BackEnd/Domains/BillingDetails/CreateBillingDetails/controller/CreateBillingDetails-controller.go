// auto-generated with ginshot
package controller

import (
	requests "create-billing-details-service/data/requests"
	responses "create-billing-details-service/data/responses"
	services "create-billing-details-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBillingDetailsController struct {
	CreateBillingDetailsService services.CreateBillingDetailsService
}

func NewCreateBillingDetailsController(service services.CreateBillingDetailsService) *CreateBillingDetailsController {
	return &CreateBillingDetailsController{
		CreateBillingDetailsService: service,
	}
}

func (ctrl *CreateBillingDetailsController) CreateBillingDetails(c *gin.Context) {
	var request requests.Request
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.Response{Message: "Invalid request body"})
		return
	}
	status, res := ctrl.CreateBillingDetailsService.CreateBillingDetailsHandler(request)

	c.IndentedJSON(status, res)
}
