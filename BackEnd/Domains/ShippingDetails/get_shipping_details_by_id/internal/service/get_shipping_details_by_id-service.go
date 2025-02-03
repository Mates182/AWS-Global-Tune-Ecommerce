// auto-generated with ginshot
package service

import (
	requests "get_shipping_details_by_id/internal/data/requests"
	responses "get_shipping_details_by_id/internal/data/responses"
)

type GetShippingDetailsByIdService interface {
	GetShippingDetailsByIdHandler(request requests.GetShippingDetailsByIdRequest) (int, responses.GetShippingDetailsByIdResponse)
}
