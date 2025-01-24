package controllers

import (
	"net/http"

	"login-service/data/request"
	"login-service/data/response"
	"login-service/service"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginService service.LoginService
}

func NewLoginController(service service.LoginService) *LoginController {
	return &LoginController{
		LoginService: service,
	}
}

func (controller *LoginController) PostLogin(c *gin.Context) {
	var loginReq request.Request
	if err := c.BindJSON(&loginReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, response.Response{Message: "Invalid request body"})
		return
	}
	status, res, cookie := controller.LoginService.LoginUser(loginReq)

	if cookie != nil {
		http.SetCookie(c.Writer, cookie)
	}

	c.IndentedJSON(status, res)
}
