package response

import "create-order-service/model"

type Response struct {
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	Order model.Order `json:"order"`
}
