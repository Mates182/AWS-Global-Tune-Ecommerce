package request

import "get-shipping-details-by-id/models"

type CreateShippingDetailsRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
