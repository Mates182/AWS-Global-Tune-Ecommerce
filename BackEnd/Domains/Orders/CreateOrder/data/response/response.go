package response

import "create-order-service/model"

type Response struct {
	Message string `json:"message"`
	Data    struct {
		Order model.Order `json:"order"`
	} `json:"data"`
}
