package request

import "update_shipping_details/internal/data/models"

type UpdateShippingDetailsRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
