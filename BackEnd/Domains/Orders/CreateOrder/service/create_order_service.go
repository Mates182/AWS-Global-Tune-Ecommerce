package service

import (
	"create-order-service/data/request"
	"create-order-service/data/response"
)

type CreateOrderService interface {
	CreateOrder(request request.Request) (int, response.Response)
}
