package service

import (
	"login-service/business-logic"
	"login-service/data/request"
	"login-service/data/response"
	"net/http"
)

type LoginServiceImpl struct {
	JWTKey []byte
}

func NewLoginServiceImpl(jwtKey []byte) LoginService {
	return &LoginServiceImpl{
		JWTKey: jwtKey,
	}
}

func (service *LoginServiceImpl) LoginUser(credentials request.Request) (int, response.Response, *http.Cookie) {
	// TODO: Validate with real data
	if credentials.Email == "admin@admin" && credentials.Password == "admin" {

		userID := "12345"
		role := "admin"
		jwtKey := service.JWTKey
		token, err := business.CreateToken(credentials.Email, userID, role, jwtKey)
		if err != nil {
			return http.StatusInternalServerError, response.Response{Message: "Could not generate token"}, nil
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

		return http.StatusOK, response.Response{Message: "Login successful"}, cookie

	}
	return http.StatusUnauthorized, response.Response{Message: "Invalid credentials"}, nil
}
