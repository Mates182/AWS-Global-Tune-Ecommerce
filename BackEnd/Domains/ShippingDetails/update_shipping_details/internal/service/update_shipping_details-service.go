// auto-generated with ginshot
package service

import (
	requests "update_shipping_details/internal/data/requests"
	responses "update_shipping_details/internal/data/responses"
)

type UpdateShippingDetailsService interface {
	UpdateShippingDetailsHandler(request requests.UpdateShippingDetailsRequest) (int, responses.UpdateShippingDetailsResponse)
}
