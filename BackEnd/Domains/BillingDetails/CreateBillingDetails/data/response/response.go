package response

import "create-billing-details-service/model"

type Response struct {
	Message string       `json:"message"`
	Data    ResponseData `json:"data"`
}

type ResponseData struct {
	BillingDetails model.BillingDetails `json:"billing_details"`
}
