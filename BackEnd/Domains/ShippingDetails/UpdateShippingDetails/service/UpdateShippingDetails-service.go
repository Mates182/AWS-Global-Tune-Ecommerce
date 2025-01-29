// auto-generated with ginshot
package service

import (
	requests "update-shipping-details/data/requests"
	responses "update-shipping-details/data/responses"
)

type UpdateShippingDetailsService interface {
	UpdateShippingDetailsHandler(request requests.UpdateShippingDetailsRequest) (int, responses.UpdateShippingDetailsResponse)
}
