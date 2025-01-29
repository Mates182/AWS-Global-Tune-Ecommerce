package request

import "create-shipping-details/models"

type CreateShippingDetailsRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
