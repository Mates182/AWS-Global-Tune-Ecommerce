package response

import "list_shipping_details/internal/data/models"

type ListShippingDetailsResponse struct {
	Message string `json:"Message"`
	ShippingDetails []models.ShippingDetails `json:"ShippingDetails"`
}
