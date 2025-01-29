package response

import "get-shipping-details-by-id/models"

type GetShippingDetailsByIdResponse struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Message string `json:"Message"`
}
