package response

import "create_shipping_details/internal/data/models"

type CreateShippingDetailsResponse struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Message string `json:"Message"`
}
