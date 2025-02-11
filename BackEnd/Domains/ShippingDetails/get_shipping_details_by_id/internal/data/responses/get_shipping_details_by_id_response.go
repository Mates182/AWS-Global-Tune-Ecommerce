package response

import "get_shipping_details_by_id/internal/data/models"

type GetShippingDetailsByIdResponse struct {
	ShippingDetails models.ShippingDetails `json:"ShippingDetails"`
	Message string `json:"Message"`
}
