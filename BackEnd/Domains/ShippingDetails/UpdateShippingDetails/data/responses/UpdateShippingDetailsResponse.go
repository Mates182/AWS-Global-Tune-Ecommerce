package response

import "update-shipping-details/models"

type UpdateShippingDetailsResponse struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Message string `json:"Message"`
}
