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

type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

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

	router.Run("localhost:8082")
}

func postLogin(c *gin.Context) {
	var login Login
	if err := c.BindJSON(&login); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	// TODO: Validate with real data
	if login.Email == "admin@admin" && login.Password == "admin" {

		userID := "12345"
		role := "admin"

		token, err := createToken(login.Email, userID, role)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
			return
		}

		cookie := &http.Cookie{
			Name:     "token",
			Value:    token,
			Path:     "/",
			Domain:   "localhost",
			HttpOnly: true,
			Secure:   false, // Change to true on production
			MaxAge:   1,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(c.Writer, cookie)

		c.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful"})
		return
	}

	c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
}

func createToken(email, userID, role string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
		Role:   role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
