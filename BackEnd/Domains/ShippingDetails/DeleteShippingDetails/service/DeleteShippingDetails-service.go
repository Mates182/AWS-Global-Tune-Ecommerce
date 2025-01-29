// auto-generated with ginshot
package service

import (
	requests "delete-shipping-details/data/requests"
	responses "delete-shipping-details/data/responses"
)

type DeleteShippingDetailsService interface {
	DeleteShippingDetailsHandler(request requests.DeleteShippingDetailsRequest) (int, responses.DeleteShippingDetailsResponse)
}
