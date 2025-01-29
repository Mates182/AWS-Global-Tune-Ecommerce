// auto-generated with ginshot
package service

import (
	requests "get-shipping-details-by-id/data/requests"
	responses "get-shipping-details-by-id/data/responses"
)

type GetShippingDetailsByIdService interface {
	GetShippingDetailsByIdHandler(request requests.GetShippingDetailsByIdRequest) (int, responses.GetShippingDetailsByIdResponse)
}
