package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtKey = []byte("my_secret_key")

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// RESTful routes
	router.POST("login", postLogin)
	router.POST("logout", postLogout)

	router.Run("localhost:8082")
}

func postLogin(c *gin.Context) {
	var login Login
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	if login.Email == "admin@admin" && login.Password == "admin" {

		token, err := createToken(login.Email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
			return
		}

		c.SetCookie("token", token, 3600, "/", "localhost", false, true)

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful"})
		return
	}

	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
}

func createToken(email string) (string, error) {

	claims := &jwt.RegisteredClaims{
		Subject:   email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func postLogout(c *gin.Context) {
	tokenString, err := c.Cookie("myTokenName")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No cookie found"})
		return
	}

	claims := &jwt.RegisteredClaims{}
	_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	c.SetCookie("myTokenName", "", 0, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
