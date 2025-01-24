package service

import (
	"logout-service/data/request"
	"logout-service/data/response"
	"net/http"
)

type LogoutService interface {
	LogoutUser(logoutReq request.Request) (int, response.Response, *http.Cookie)
}
