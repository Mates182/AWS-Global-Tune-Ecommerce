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

// PostLogin 	godoc
// @Summary 	Authenticate user and issue a JWT token
// @Description Validates user credentials and returns a JWT token in a cookie upon successful authentication.
// @Tags 		Authentication
// @Accept 		json
// @Produce 	json
// @Param 		login body request.Request true "User login credentials"
// @Success 	200 {object} response.Response "Login successful"
// @Failure 	400 {object} response.Response "Invalid request body"
// @Failure 	401 {object} response.Response "Invalid credentials"
// @Failure 	500 {object} response.Response "Could not generate token"
// @Router 		/login [post]
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
