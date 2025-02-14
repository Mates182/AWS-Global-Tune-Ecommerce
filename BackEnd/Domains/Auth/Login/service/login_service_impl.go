package service

import (
	"database/sql"
	"log"
	"login-service/business-logic"
	"login-service/data/request"
	"login-service/data/response"
	"net/http"
)

type LoginServiceImpl struct {
	JWTKey []byte
	DB     *sql.DB
}

func NewLoginServiceImpl(jwtKey []byte, db *sql.DB) LoginService {
	return &LoginServiceImpl{
		JWTKey: jwtKey,
		DB:     db,
	}
}

func (service *LoginServiceImpl) LoginUser(credentials request.Request) (int, response.Response, *http.Cookie) {
	var userID, name, email, role string

	query := "CALL GetUserByEmailAndPassword(?, ?)"
	err := service.DB.QueryRow(query, credentials.Email, credentials.Password).Scan(&userID, &name, &email, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusUnauthorized, response.Response{Message: "Invalid credentials"}, nil
		}
		log.Printf("Database error: %v", err)
		return http.StatusInternalServerError, response.Response{Message: "Database error"}, nil
	}

	token, err := business.CreateToken(email, userID, role, service.JWTKey)
	if err != nil {
		return http.StatusInternalServerError, response.Response{Message: "Could not generate token"}, nil
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "52.90.4.29",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   3600,
		SameSite: http.SameSiteNoneMode,
	}

	return http.StatusOK, response.Response{Message: "Login successful"}, cookie
}
