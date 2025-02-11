package response

import "update_shipping_details/internal/data/models"

type UpdateShippingDetailsResponse struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Message string `json:"Message"`
}
