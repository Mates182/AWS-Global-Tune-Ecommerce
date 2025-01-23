package controllers

import (
	"net/http"

	"login-service/data/request"
	"login-service/data/response"
	"login-service/secrets"

	"github.com/gin-gonic/gin"
)

func PostLogin(c *gin.Context) {
	var login request.Request
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, response.Response{Message: "Invalid request body"})
		return
	}
	// TODO: Validate with real data
	if login.Email == "admin@admin" && login.Password == "admin" {

		userID := "12345"
		role := "admin"
		jwtKey := secrets.GetJWTKey()
		token, err := CreateToken(login.Email, userID, role, jwtKey)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, response.Response{Message: "Could not generate token"})
			return
		}

		cookie := &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: true,
			Secure:   false, // Change to true on production
			MaxAge:   3600,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(c.Writer, cookie)

		c.IndentedJSON(http.StatusOK, response.Response{Message: "Login successful"})
		return
	}

	c.IndentedJSON(http.StatusUnauthorized, response.Response{Message: "Invalid credentials"})
}
