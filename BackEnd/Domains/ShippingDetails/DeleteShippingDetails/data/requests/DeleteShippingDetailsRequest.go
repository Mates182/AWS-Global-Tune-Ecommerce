package request

import "delete-shipping-details/models"

type DeleteShippingDetailsRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
