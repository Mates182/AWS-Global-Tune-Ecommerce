package controllers

import (
	"net/http"

	"logout-service/data/request"
	"logout-service/service"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutService service.LogoutService
}

func NewLogoutController(service service.LogoutService) *LogoutController {
	return &LogoutController{
		LogoutService: service,
	}
}

func (controller *LogoutController) PostLogout(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No cookie found"})
		return
	}
	logoutRequest := &request.Request{Token: tokenString}
	status, res, cookie := controller.LogoutService.LogoutUser(*logoutRequest)

	if cookie != nil {
		http.SetCookie(c.Writer, cookie)
	}

	c.IndentedJSON(status, res)
}
