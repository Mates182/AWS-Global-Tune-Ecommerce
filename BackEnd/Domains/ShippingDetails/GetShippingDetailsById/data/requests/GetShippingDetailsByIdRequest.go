package request

import "get-shipping-details-by-id/models"

type GetShippingDetailsByIdRequest struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
}
