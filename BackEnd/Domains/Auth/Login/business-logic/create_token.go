package business

import (
	"login-service/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(email, userID, role string, jwtKey []byte) (string, error) {
	claims := &models.Claims{
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
