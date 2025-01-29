package response

import "delete-shipping-details/models"

type DeleteShippingDetailsResponse struct {
	Message string `json:"Message"`
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
