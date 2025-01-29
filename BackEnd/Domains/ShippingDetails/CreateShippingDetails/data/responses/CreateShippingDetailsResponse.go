package response

import "create-shipping-details/models"

type CreateShippingDetailsResponse struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Message string `json:"Message"`
}
