package request

import "create-billing-details-service/model"

type Request struct {
	Data struct {
		BillingDetails model.BillingDetails `json: "billing_details" bson: "billing_details"`
	} `json:"data"`
}
