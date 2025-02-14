package service

import (
	"logout-service/data/request"
	"logout-service/data/response"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type LogoutServiceImpl struct {
	JWTKey []byte
}

func NewLogoutServiceImpl(jwtKey []byte) LogoutService {
	return &LogoutServiceImpl{
		JWTKey: jwtKey,
	}
}

func (service *LogoutServiceImpl) LogoutUser(logoutReq request.Request) (int, response.Response, *http.Cookie) {

	jwtKey := service.JWTKey
	tokenString := logoutReq.Token
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return http.StatusUnauthorized, response.Response{Message: "Invalid token"}, nil
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   "http://52.90.4.29",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
	}

	return http.StatusOK, response.Response{Message: "Logged out successfully"}, cookie

}
