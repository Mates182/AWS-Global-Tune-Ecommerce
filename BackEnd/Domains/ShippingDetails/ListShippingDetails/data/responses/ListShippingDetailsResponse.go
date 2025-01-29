package response

import "list-shipping-details/models"

type ListShippingDetailsResponse struct {
	ShippingDetails []models.ShippingDetails `json:"ShippingDetails"`
	Message         string                   `json:"Message"`
}
