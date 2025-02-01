package service

import (
	"login-service/data/request"
	"login-service/data/response"
	"net/http"
)

type LoginService interface {
	LoginUser(credentials request.Request) (int, response.Response, *http.Cookie)
}
