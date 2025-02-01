package request

import "update-shipping-details/models"

type UpdateShippingDetailsRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
