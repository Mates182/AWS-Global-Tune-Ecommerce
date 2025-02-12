package request

import "create_shipping_details/internal/data/models"

type CreateShippingDetailsRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
