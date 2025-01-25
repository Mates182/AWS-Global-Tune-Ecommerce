package controllers

import (
	"net/http"

	roleAuth "create-order-service/api/role-auth"
	"create-order-service/data/request"
	"create-order-service/data/response"
	"create-order-service/service"

	"github.com/gin-gonic/gin"
)

type CreateOrderController struct {
	CreateOrderService service.CreateOrderService
}

func NewCreateOrderController(service service.CreateOrderService) *CreateOrderController {
	return &CreateOrderController{
		CreateOrderService: service,
	}
}

func (controller *CreateOrderController) CreateOrder(c *gin.Context) {
	var orderReq request.Request
	if err := c.BindJSON(&orderReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, response.Response{Message: "Invalid request body"})
		return
	}
	token, err := c.Cookie("token")
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, response.Response{Message: "Missing or invalid token"})
		return
	}
	if !roleAuth.AuthenticateUser(token) {
		c.IndentedJSON(http.StatusUnauthorized, response.Response{Message: "Unauthorized"})
		return

	}
	status, res := controller.CreateOrderService.CreateOrder(orderReq)

	c.IndentedJSON(status, res)
}
